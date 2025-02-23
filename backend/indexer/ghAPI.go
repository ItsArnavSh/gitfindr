package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func fetchJSON(url, token string, result any) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	if token != "" {
		req.Header.Set("Authorization", "token "+token)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body) // Reading response body for error details
		return fmt.Errorf("GitHub API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	return nil
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
		return nil, errors.New("invalid repository URL format")
	}
	owner, repo := parts[len(parts)-2], parts[len(parts)-1]
	baseURL := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)

	var repoData map[string]any
	if err := fetchJSON(baseURL, token, &repoData); err != nil {
		return nil, fmt.Errorf("failed to fetch repository info: %w", err)
	}

	getString := func(data map[string]any, key string) string {
		if val, ok := data[key].(string); ok {
			return val
		}
		return ""
	}

	getInt := func(data map[string]any, key string) int {
		if val, ok := data[key].(float64); ok {
			return int(val)
		}
		return 0
	}

	var topics []string
	if t, ok := repoData["topics"].([]any); ok {
		for _, topic := range t {
			if str, ok := topic.(string); ok {
				topics = append(topics, str)
			}
		}
	}

	var languages map[string]any
	if err := fetchJSON(baseURL+"/languages", token, &languages); err != nil {
		return nil, fmt.Errorf("failed to fetch languages: %w", err)
	}
	var languageList []string
	for lang := range languages {
		languageList = append(languageList, lang)
	}

	var tags []map[string]any
	if err := fetchJSON(baseURL+"/tags", token, &tags); err != nil {
		return nil, fmt.Errorf("failed to fetch tags: %w", err)
	}
	var tagList []string
	for _, tag := range tags {
		if name, ok := tag["name"].(string); ok {
			tagList = append(tagList, name)
		}
	}

	var contributors []map[string]any
	if err := fetchJSON(baseURL+"/contributors", token, &contributors); err != nil {
		return nil, fmt.Errorf("failed to fetch contributors: %w", err)
	}
	var contributorList []string
	for _, contributor := range contributors {
		if login, ok := contributor["login"].(string); ok {
			contributorList = append(contributorList, login)
		}
	}

	var commits []map[string]any
	if err := fetchJSON(baseURL+"/commits", token, &commits); err != nil {
		return nil, fmt.Errorf("failed to fetch commits: %w", err)
	}
	lastCommitDate := ""
	if len(commits) > 0 {
		if commit, ok := commits[0]["commit"].(map[string]any); ok {
			if committer, ok := commit["committer"].(map[string]any); ok {
				lastCommitDate = getString(committer, "date")
			}
		}
	}

	licenseName := "No License"
	if license, ok := repoData["license"].(map[string]any); ok {
		licenseName = getString(license, "name")
	}

	repoInfo := &RepoInfo{
		Name:         getString(repoData, "name"),
		Owner:        getString(repoData["owner"].(map[string]any), "login"),
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

	repoInfoJSON, err := json.Marshal(repoInfo)
	if err == nil {
		AddToCache(repoURL, string(repoInfoJSON))
	}

	return repoInfo, nil
}
