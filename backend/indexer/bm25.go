package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	_ "github.com/glebarez/go-sqlite"
)

var k float64 = 1.5

func bm25(query string) {
	normalizedQuery := parseWords(query)
	queryTermIDs := make(map[int]bool) // Unique query term IDs
	var indexes []int

	// Step 1: Get index list for all query terms
	for _, token := range normalizedQuery {
		res, _ := rdb.Exists(ctx, token).Result()
		if res > 0 {
			vals, _ := FetchList(token)
			fmt.Printf("\nToken: '%s' → TermIDs: %v\n", token, vals)
			for _, val := range vals {
				queryTermIDs[val] = true
				indexes = append(indexes, val)
			}
		} else {
			fmt.Printf("\nToken: '%s' → ❌ No matching TermID found\n", token)
		}
	}

	queryTermCount := len(queryTermIDs)
	fmt.Println("\nFinal Unique TermIDs Used:", queryTermIDs)

	// Step 2: BM25 Calculation
	bmScores := make(map[int]float64)        // Mapping Doc ID to BM25 score
	docTermSet := make(map[int]map[int]bool) // Tracks query TermIDs present in each doc
	docDetails := make(map[int][]string)     // Stores BM25 breakdown per document
	fmt.Println("\n===== SEARCH PROCESSING =====")

	for _, termID := range indexes {
		rowData, err := getInvertedIndex(termID)
		emRowData, emErr := getemInvertedIndex(termID) // Fetch extra BM25 index

		if err != nil {
			fmt.Println("Error Querying getInvertedIndex")
			log.Fatal(err)
		}
		if emErr != nil {
			fmt.Println("Error Querying getemInvertedIndex")
			log.Fatal(emErr)
		}

		// Combine both sources
		processBM25Data(rowData, bmScores, docTermSet, docDetails, termID, 1.0)   // Normal weight
		processBM25Data(emRowData, bmScores, docTermSet, docDetails, termID, 2.0) // Double weight
	}

	// Step 3: Apply a boost if ALL query term IDs exist in the document
	for docID := range bmScores {
		if len(docTermSet[docID]) == queryTermCount {
			bmScores[docID] *= 1 // Adjust boost factor if needed
			docDetails[docID] = append(docDetails[docID], fmt.Sprintf("Boost applied (all terms present) | New Score: %.4f", bmScores[docID]))
		}
	}

	// Step 4: Sort Documents by BM25 Score
	sortedDocs := sortDocsByBM25(bmScores)

	// Step 5: Limit to Top 25
	if len(sortedDocs) > 5 {
		sortedDocs = sortedDocs[:5]
	}

	// Step 6: Fetch and Display Document Details
	fmt.Println("\n===== BM25 DEBUG INFO =====")
	for rank, docID := range sortedDocs {
		fmt.Printf("\nRank #%d | DocID: %d | Final BM25 Score: %.4f\n", rank+1, docID, bmScores[docID])
		fmt.Println("Breakdown:")
		for _, detail := range docDetails[docID] {
			fmt.Println("  ", detail)
		}
	}

	// Display final results
	displayResults(sortedDocs)
}

// Helper function to process BM25 data with weighting
func processBM25Data(rowData *InvertedIndex, bmScores map[int]float64, docTermSet map[int]map[int]bool, docDetails map[int][]string, termID int, weight float64) {
	for _, repo := range rowData.ReposWeighUnweigh {
		if _, exists := bmScores[repo.Repo]; !exists {
			bmScores[repo.Repo] = 0.0
			docTermSet[repo.Repo] = make(map[int]bool)
		}

		// BM25 Parameters
		b := 0.9
		k := 1.2
		doc, _ := getDoc(repo.Repo)
		avgDL, _ := getAverageDocLength()

		weightedTermFrequency := float64(repo.Weighted)
		idf := rowData.IDF
		if idf < 0.1 {
			idf = 0.5
		}
		docLength := float64(doc.DocLength)
		termScore := weight * idf * ((weightedTermFrequency * (k + 1)) /
			(weightedTermFrequency + k*(1-b+b*(docLength/avgDL))))

		bmScores[repo.Repo] += termScore
		docTermSet[repo.Repo][termID] = true // Track presence of TermID

		// Store breakdown for debugging
		detail := fmt.Sprintf("TermID: %d | IDF: %.4f | WeightedTF: %.2f | DocLen: %.2f | AvgDL: %.2f | Weight: %.1f | Score: %.4f",
			termID, idf, weightedTermFrequency, docLength, avgDL, weight, termScore)

		docDetails[repo.Repo] = append(docDetails[repo.Repo], detail)
	}
}

func getAverageDocLength() (float64, error) {
	query := `SELECT AVG(docLength) FROM documents`
	row := db.QueryRow(query)

	var avgLength float64

	// Scan into the variable
	err := row.Scan(&avgLength)
	if err != nil {
		return 0, err
	}

	return avgLength, nil
}

func getDoc(docID int) (*Document, error) {
	query := `SELECT docID, name, link, alt, docLength FROM documents WHERE docID = ?`
	row := db.QueryRow(query, docID)

	var doc Document

	// Scan into struct fields
	err := row.Scan(&doc.DocID, &doc.Name, &doc.Link, &doc.Alt, &doc.DocLength)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

func getInvertedIndex(termID int) (*InvertedIndex, error) {
	query := `SELECT termID, reposWeighUnweigh, dfi, IDF FROM invertedIndex WHERE termID = ?`
	row := db.QueryRow(query, termID)

	var index InvertedIndex
	var jsonStr string

	// Scan into variables
	err := row.Scan(&index.TermID, &jsonStr, &index.DFI, &index.IDF)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON
	err = json.Unmarshal([]byte(jsonStr), &index.ReposWeighUnweigh)
	if err != nil {
		return nil, err
	}

	return &index, nil
}
func getemInvertedIndex(termID int) (*InvertedIndex, error) {
	query := `SELECT termID, reposWeighUnweigh, dfi, IDF FROM emInvertedIndex WHERE termID = ?`
	row := db.QueryRow(query, termID)

	var index InvertedIndex
	var jsonStr string

	// Scan into variables
	err := row.Scan(&index.TermID, &jsonStr, &index.DFI, &index.IDF)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON
	err = json.Unmarshal([]byte(jsonStr), &index.ReposWeighUnweigh)
	if err != nil {
		return nil, err
	}

	return &index, nil
}

func displayResults(sortedDocIDs []int) {
	fmt.Println("\nTop Results:")
	for _, docID := range sortedDocIDs {
		name, link, err := getDocument(docID)
		if err != nil {
			fmt.Println("Error fetching document:", err)
			continue
		}
		fmt.Printf("DocID: %d | Name: %s | Link: %s\n", docID, name, link)
	}
}
func getDocument(docID int) (string, string, error) {
	query := `SELECT name, link FROM documents WHERE docID = ?`
	row := db.QueryRow(query, docID)

	var name, link string
	err := row.Scan(&name, &link)
	if err != nil {
		return "", "", err
	}
	return name, link, nil
}

func sortDocsByBM25(bmScores map[int]float64) []int {
	type docScore struct {
		DocID int
		Score float64
	}

	// Convert map to slice for sorting
	var docList []docScore
	for docID, score := range bmScores {
		docList = append(docList, docScore{DocID: docID, Score: score})
	}

	// Sort by Score in Descending Order
	sort.Slice(docList, func(i, j int) bool {
		return docList[i].Score > docList[j].Score
	})

	// Extract sorted Doc IDs
	var sortedDocIDs []int
	for _, doc := range docList {
		sortedDocIDs = append(sortedDocIDs, doc.DocID)
	}
	return sortedDocIDs
}
