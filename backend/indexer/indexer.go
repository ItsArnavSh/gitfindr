package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"

	_ "github.com/glebarez/go-sqlite"
)

func index(link string) {
	word_list, err := getRepoData(link)
	if err != nil {
		return
	}
	indexList := goClusterer(word_list)
	emIndexList := EMCluster(word_list)
	repoInfo, err := getRepoInfo(link)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	indexer(indexList, emIndexList, *repoInfo, link)
	fmt.Println("Done ", repoInfo.Name)
}
func indexer(indexList []int, emIndexList []int, repoInfo RepoInfo, link string) {
	docLen := len(indexList)
	starConstant, forkConstant := 5.0, 10.0
	alt := 1 + starConstant*math.Log(float64(repoInfo.Stars)+1) + forkConstant*math.Log(float64(repoInfo.Forks)+1)
	insertQuery := `INSERT INTO documents (name, link, alt, docLength) VALUES (?, ?, ?, ?);`
	result, err := db.Exec(insertQuery, repoInfo.Name, link, alt, docLen)
	if err != nil {
		log.Fatal(err)
	}
	docID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	var weigher []wordWeigher
	var emweigher []wordWeigher
	desc := parseWords(repoInfo.Description)
	descList := freqMap(goClusterer(desc))
	emdescList := freqMap(EMCluster(desc))
	weigher = append(weigher, wordWeigher{emdescList, 5})
	emweigher = append(emweigher, wordWeigher{descList, 5})

	tags := repoInfo.Tags
	tagIndex := freqMap(goClusterer(tags))
	emtagIndex := freqMap(EMCluster(tags))
	weigher = append(weigher, wordWeigher{tagIndex, 4}) // Priority for tags
	emweigher = append(emweigher, wordWeigher{emtagIndex, 4})

	lang := repoInfo.Languages
	langIndex := freqMap(goClusterer(lang))
	emlangIndex := freqMap(EMCluster(lang))
	weigher = append(weigher, wordWeigher{langIndex, 5}) // Priority for languages
	emweigher = append(emweigher, wordWeigher{emlangIndex, 5})

	topics := repoInfo.Topics
	topicIndex := freqMap(goClusterer(topics))
	emtopicIndex := freqMap(EMCluster(topics))
	weigher = append(weigher, wordWeigher{topicIndex, 3}) // Priority for topics
	emweigher = append(emweigher, wordWeigher{emtopicIndex, 3})

	bodyFreq := freqMap(indexList)
	embodyFreq := freqMap(emIndexList)
	weigher = append(weigher, wordWeigher{bodyFreq, 1}) // Priority for body
	emweigher = append(emweigher, wordWeigher{embodyFreq, 1})

	data := freqCombiner(weigher)
	emdata := freqCombiner(emweigher)

	freqMapToII(db, data.unweighted, data.weighted, int(docID), false)
	freqMapToII(db, emdata.unweighted, emdata.weighted, int(docID), true)

	fmt.Println(docID)
}

// freqMap converts an integer array into a map[int]int where keys are numbers and values are their frequencies
func freqMap(arr []int) map[int]int {
	freq := make(map[int]int)

	for _, num := range arr {
		freq[num]++
	}

	return freq
}

func freqMapToII(db *sql.DB, freqMap map[int]int, weightdeFreqMap map[int]int, docID int, em bool) {
	table := "invertedIndex"
	if !em {
		table = "emInvertedIndex"
	}

	for termID, freq := range freqMap {
		var existingJSON string
		var dfi int
		// Check if termID exists in the table
		err := db.QueryRow("SELECT reposWeighUnweigh, dfi FROM "+table+" WHERE termID = ?", termID).Scan(&existingJSON, &dfi)
		if err != nil {
			if err == sql.ErrNoRows {
				// If term does not exist, create a new JSON array with the first element
				newEntry := []reposWeighUnweigh{
					{Weighted: weightdeFreqMap[termID], Unweighted: freq, Repo: docID},
				}
				jsonData, _ := json.Marshal(newEntry)
				dfi := len(newEntry)
				idf := math.Log((float64(docID-dfi)+float64(0.5))/(float64(dfi)+float64(0.5)) + 1)
				_, err = db.Exec("INSERT INTO "+table+" (termID, reposWeighUnweigh, dfi, IDF) VALUES (?, ?, ?, ?)",
					termID, jsonData, len(newEntry), idf)
				if err != nil {
					log.Println("Error inserting new term:", err)
				}
			} else {
				log.Println("Error querying termID:", err)
			}
			continue
		}

		// If term exists, append the new repo entry
		var existingEntries []reposWeighUnweigh
		err = json.Unmarshal([]byte(existingJSON), &existingEntries)
		if err != nil {
			log.Println("Error unmarshalling JSON:", err)
			continue
		}

		existingEntries = append(existingEntries, reposWeighUnweigh{Weighted: weightdeFreqMap[termID], Unweighted: freq, Repo: docID})
		updatedJSON, _ := json.Marshal(existingEntries)

		// Update the table with the new JSON and increment dfi
		_, err = db.Exec("UPDATE "+table+" SET reposWeighUnweigh = ?, dfi = ? WHERE termID = ?", updatedJSON, dfi+1, termID)
		if err != nil {
			log.Println("Error updating termID:", err)
		}
	}
}

func arrToTechWords(terms []string) []Word {
	var techWords []Word
	for _, term := range terms {
		techWords = append(techWords, Word{term, "tech term"})
	}
	return techWords
}

func freqCombiner(weights []wordWeigher) weighUnweigh {
	//Combines them
	var data weighUnweigh
	data.weighted = make(map[int]int)
	data.unweighted = make(map[int]int)
	for _, wordInfo := range weights {
		for words := range wordInfo.freqMap {
			_, ok := data.weighted[words]
			if ok {
				data.weighted[words] += wordInfo.freqMap[words] * wordInfo.weight
				data.unweighted[words] += wordInfo.freqMap[words]
			} else {
				data.weighted[words] = wordInfo.freqMap[words] * wordInfo.weight
				data.unweighted[words] = wordInfo.freqMap[words]
			}
		}
	}
	return data
}
