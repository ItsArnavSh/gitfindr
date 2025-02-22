package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/glebarez/go-sqlite"
	"github.com/redis/go-redis/v9"
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

// AppendIntToRedisList appends a given integer to a Redis list stored under the given key.
func AppendIntToRedisList(key string, num int) error {
	// Convert integer to string
	numStr := strconv.Itoa(num)

	// Append to the list (RPUSH adds to the end)
	err := rdb.RPush(ctx, key, numStr).Err()
	if err != nil {
		return fmt.Errorf("failed to append to Redis list: %v", err)
	}

	return nil
}

// FetchList fetches all values from the Redis list and converts them back to integers.
func FetchList(key string) ([]int, error) {
	// Get the full list from Redis
	values, err := rdb.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Redis list: %v", err)
	}

	// Convert string values to integers
	var result []int
	for _, v := range values {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("error converting value to int: %v", err)
		}
		result = append(result, num)
	}

	return result, nil
}

const (
	mappingKey = "string_to_number" // Redis hash key for mapping
	counterKey = "string_counter"   // Redis key for global counter
)

func SetStringGetNumber(s string) (int, error) {
	// Check if the string already has a number assigned
	num, err := rdb.HGet(ctx, mappingKey, s).Int64()
	if err == nil {
		return int(num), nil // Return existing mapping
	}
	if err != redis.Nil {
		return 0, err // Return error if it's not a "key does not exist" case
	}

	// Get next number
	num, err = rdb.Incr(ctx, counterKey).Result()
	if err != nil {
		return 0, err
	}

	// Store the mapping
	if err := rdb.HSet(ctx, mappingKey, s, num).Err(); err != nil {
		return 0, err
	}

	return int(num), nil
}
