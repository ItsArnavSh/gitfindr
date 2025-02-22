package main

type Word struct {
	Term     string `json:"word"`
	TermType string `json:"type"`
}

type WordEntry struct {
	Word string `json:"word"`
}

type RepoInfo struct {
	Name         string   `json:"name"`
	Owner        string   `json:"owner"`
	Description  string   `json:"description"`
	Stars        int      `json:"stargazers_count"`
	Forks        int      `json:"forks_count"`
	OpenIssues   int      `json:"open_issues_count"`
	License      string   `json:"license"`
	Topics       []string `json:"topics"`
	Languages    []string `json:"languages"`
	Tags         []string `json:"tags"`
	Contributors []string `json:"contributors"`
	TotalCommits int      `json:"total_commits"`
	LastCommit   string   `json:"last_commit_date"`
}

//Used in indexing

type wordWeigher struct {
	freqMap map[int]int
	weight  int
}
type weighUnweigh struct {
	weighted   map[int]int
	unweighted map[int]int
}
type reposWeighUnweigh struct {
	Weighted   int `json:"weighted"`
	Unweighted int `json:"unweighted"`
	Repo       int `json:"repoid"`
}

// For Bm25
type InvertedIndex struct {
	TermID            int                 `json:"termID"`
	ReposWeighUnweigh []reposWeighUnweigh `json:"reposWeighUnweigh"`
	DFI               int                 `json:"dfi"`
	IDF               float64             `json:"idf"`
}
type Document struct {
	DocID     int
	Name      string
	Link      string
	Alt       float64
	DocLength int
}
