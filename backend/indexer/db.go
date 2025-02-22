package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

func connectSQL() {
	db, _ = sql.Open("sqlite", "./testDB.db")
	db.Ping()
	fmt.Println("Connected to the SQLite database successfully.")
}

func setUpSQLITE() {
	query := `CREATE TABLE IF NOT EXISTS documents(
	docID INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	link TEXT,
	alt REAL,
	docLength INTEGER
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	query = `CREATE TABLE IF NOT EXISTS invertedIndex(
	termID INTEGER PRIMARY KEY,
	reposWeighUnweigh TEXT,
	dfi INTEGER,
	IDF REAL
	);
	`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	query = `CREATE TABLE IF NOT EXISTS emInvertedIndex(
	termID INTEGER PRIMARY KEY,
	reposWeighUnweigh TEXT,
	dfi INTEGER,
	IDF REAL
	);
	`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created or already exists.")
}
func InitCache() { //For Caching purposes
	query := `CREATE TABLE IF NOT EXISTS apiCache (
		payload TEXT PRIMARY KEY,
		response TEXT
	)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create cache table: %v", err)
	}
}
