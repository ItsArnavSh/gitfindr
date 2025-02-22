package main

type Word struct {
	Term     string `json:"word"`
	TermType string `json:"type"`
}

type WordEntry struct {
	Word string `json:"word"`
}
