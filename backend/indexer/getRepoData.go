package main

import (
	"encoding/json"
	"fmt"

	"io"
	"log"
	"net/http"
	"strings"
)

func getRepoData(link string) ([]string, error) {

	url := "http://127.0.0.1:8000/convert"

	payload := strings.NewReader(fmt.Sprintf("{\n  \"name\":\"%s\"\n}", link))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "insomnia/10.3.0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//fmt.Println(res)
	jsonData := string(body)
	var result map[string][]string
	err = json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		return nil, err
	}
	list_words := result["data_list"]
	return list_words, nil
}

func parseWords(link string) []string {

	url := "http://127.0.0.1:8000/convertText"

	payload := strings.NewReader(fmt.Sprintf("{\n  \"text\":\"%s\"\n}", link))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "insomnia/10.3.0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//fmt.Println(res)
	jsonData := string(body)
	var result map[string][]string
	err = json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		log.Fatal(err)
	}
	list_words := result["data_list"]
	return list_words
}
