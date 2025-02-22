package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func fetchJSON(url, token string, result interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	if token != "" {
		req.Header.Set("Authorization", "token "+token)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, result)
}

func getRepoInfo(repoURL string) (*RepoInfo, error) {
	// Check cache first
	if cachedResponse, available := GetFromCache(repoURL); available {
		var repoInfo RepoInfo
		if err := json.Unmarshal([]byte(cachedResponse), &repoInfo); err == nil {
			return &repoInfo, nil
		}
	}

	token, _ := os.LookupEnv("GITHUB_API_KEY")
	// Extract owner and repo from URL
	parts := strings.Split(strings.TrimRight(repoURL, "/"), "/")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid repository URL")
	}
	owner, repo := parts[len(parts)-2], parts[len(parts)-1]
	baseURL := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)

	var repoData map[string]interface{}
	if err := fetchJSON(baseURL, token, &repoData); err != nil {
		return nil, err
	}

	// Helper functions
	getString := func(data map[string]interface{}, key string) string {
		if val, ok := data[key].(string); ok {
			return val
		}
		return ""
	}
	getInt := func(data map[string]interface{}, key string) int {
		if val, ok := data[key].(float64); ok {
			return int(val)
		}
		return 0
	}

	// Topics
	var topics []string
	if t, ok := repoData["topics"].([]interface{}); ok {
		for _, topic := range t {
			if str, ok := topic.(string); ok {
				topics = append(topics, str)
			}
		}
	}

	// Languages
	var languages map[string]interface{}
	if err := fetchJSON(baseURL+"/languages", token, &languages); err != nil {
		return nil, err
	}
	var languageList []string
	for lang := range languages {
		languageList = append(languageList, lang)
	}

	// Tags
	var tags []map[string]interface{}
	if err := fetchJSON(baseURL+"/tags", token, &tags); err != nil {
		return nil, err
	}
	var tagList []string
	for _, tag := range tags {
		if name, ok := tag["name"].(string); ok {
			tagList = append(tagList, name)
		}
	}

	// Contributors
	var contributors []map[string]interface{}
	if err := fetchJSON(baseURL+"/contributors", token, &contributors); err != nil {
		return nil, err
	}
	var contributorList []string
	for _, contributor := range contributors {
		if login, ok := contributor["login"].(string); ok {
			contributorList = append(contributorList, login)
		}
	}

	// Commits
	var commits []map[string]interface{}
	if err := fetchJSON(baseURL+"/commits", token, &commits); err != nil {
		return nil, err
	}
	lastCommitDate := ""
	if len(commits) > 0 {
		if commit, ok := commits[0]["commit"].(map[string]interface{}); ok {
			if committer, ok := commit["committer"].(map[string]interface{}); ok {
				lastCommitDate = getString(committer, "date")
			}
		}
	}

	// License
	licenseName := "No License"
	if license, ok := repoData["license"].(map[string]interface{}); ok {
		licenseName = getString(license, "name")
	}

	repoInfo := &RepoInfo{
		Name:         getString(repoData, "name"),
		Owner:        getString(repoData["owner"].(map[string]interface{}), "login"),
		Description:  getString(repoData, "description"),
		Stars:        getInt(repoData, "stargazers_count"),
		Forks:        getInt(repoData, "forks_count"),
		OpenIssues:   getInt(repoData, "open_issues_count"),
		License:      licenseName,
		Topics:       topics,
		Languages:    languageList,
		Tags:         tagList,
		Contributors: contributorList,
		TotalCommits: len(commits),
		LastCommit:   lastCommitDate,
	}

	// Store in cache
	repoInfoJSON, err := json.Marshal(repoInfo)
	if err == nil {
		AddToCache(repoURL, string(repoInfoJSON))
	}

	return repoInfo, nil
}
