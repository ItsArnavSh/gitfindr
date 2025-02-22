package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type techWord struct {
	Word string   `json:"word"`
	Syns []string `json:"synonyms"`
}

func goClusterer(terms []string) []int {

	//Process the words just in case
	for i := range terms {
		terms[i] = strings.ToLower(terms[i])
	}
	//Check if a word exists in redis
	//If it does not, if its a platform, add it directly
	//If its a normal word, ask text api,
	//If it is a tech word, then ask good gemini for synonyms
	//Get the list of synonyms and chill then
	var cluster []int
	for _, word := range terms {
		fmt.Println("Processing ", word)
		res, err := rdb.Exists(ctx, word).Result()
		if err != nil {
			log.Fatal(err)
		}
		if res > 0 {
			//fmt.Println(word, " Already Exists")
			//The term exists already
			//rawval, _ := rdb.Get(ctx, word.Term).Result()
			value, err := FetchList(word)
			//value, err = strconv.Atoi(rawval)
			if err == nil {
				for _, val := range value {
					cluster = append(cluster, val)
				}
			}
		} else {
			//Not found
			if isVersion(word) || isNumericOnly(word) {
				continue
			}
			//Now we have to see which category it belongs to and act accordingly
			rawval, _ := rdb.Get(ctx, "last_index").Result()
			lastIndex, _ := strconv.Atoi(rawval)
			//fmt.Println("Setting " + word + " Common")
			time.Sleep(200 * time.Millisecond) //Avoid Rate limit
			synonyms := getSynonyms(word)
			lastIndex++
			//rdb.Set(ctx, word.Term, lastIndex, -1)
			AppendIntToRedisList(word, lastIndex)

			for _, syn := range synonyms {
				//fmt.Println("Saving ", syn, " at ", lastIndex)
				//rdb.Set(ctx, syn, lastIndex, -1)
				AppendIntToRedisList(syn, lastIndex)
			}
			rdb.Set(ctx, "last_index", lastIndex, -1)
			cluster = append(cluster, lastIndex)
			// switch word.TermType {

			// case "non-tech term":
			// 	fmt.Println("Setting " + word.Term + " Common")
			// 	time.Sleep(200 * time.Millisecond) //Avoid Rate limit
			// 	synonyms := getSynonyms(word.Term)
			// 	lastIndex++
			// 	//rdb.Set(ctx, word.Term, lastIndex, -1)
			// 	AppendIntToRedisList(word.Term, lastIndex)

			// 	for _, syn := range synonyms {
			// 		fmt.Println("Saving ", syn, " at ", lastIndex)
			// 		//rdb.Set(ctx, syn, lastIndex, -1)
			// 		AppendIntToRedisList(syn, lastIndex)
			// 	}
			// 	rdb.Set(ctx, "last_index", lastIndex, -1)
			// 	cluster = append(cluster, lastIndex)

			// case "tech term":
			// 	fmt.Println("Setting " + word.Term + " Platform")
			// 	lastIndex++
			// 	//rdb.Set(ctx, word.Term, lastIndex, -1)
			// 	AppendIntToRedisList(word.Term, lastIndex)
			// 	rdb.Set(ctx, "last_index", lastIndex, -1)
			// 	cluster = append(cluster, lastIndex)
			// default:
			// 	continue
			// }
		}
	}

	return cluster
}

// Define a struct to match the JSON response structure

func getSynonyms(word string) []string {
	url := "https://api.onelook.com/words?ml=" + word + "&max=5"
	req, err := http.NewRequest("GET", url, nil) // Use GET instead of POST
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	// Parse the JSON response
	var response []WordEntry
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		//log.Fatalf("Error unmarshalling response: %v", err)
	}

	// Extract only the "word" values
	var words []string
	for _, entry := range response {
		words = append(words, entry.Word)
	}

	return words
}

func textToJSStr(text string) string {
	start := strings.Index(text, "{")
	end := strings.LastIndex(text, "}")

	return text[start : end+1]
}
func isVersion(s string) bool {
	re := regexp.MustCompile(`^v\d+\.\d+(\.\d+)?(-[a-zA-Z0-9]+)?$`)
	return re.MatchString(s)
}
func isNumericOnly(s string) bool {
	re := regexp.MustCompile(`^[0-9.]+$`)
	return re.MatchString(s)
}
func EMCluster(data []string) []int {
	var intData []int
	for _, word := range data {
		intD, _ := SetStringGetNumber(strings.ToLower(word))
		intData = append(intData, intD)
	}
	return intData
}
