// Simple SQLITE Cache is a table and maps api payload to response
package main

import (
	"database/sql"
	"log"
)

// GetFromCache checks if a payload exists in the cache and returns the response
func GetFromCache(payload string) (string, bool) {
	var response string
	err := db.QueryRow("SELECT response FROM apiCache WHERE payload = ?", payload).Scan(&response)
	if err == sql.ErrNoRows {
		return "", false
	} else if err != nil {
		log.Printf("Error querying cache: %v", err)
		return "", false
	}
	return response, true
}

// AddToCache inserts a new payload-response pair into the cache
func AddToCache(payload, response string) {
	_, err := db.Exec("INSERT OR REPLACE INTO apiCache (payload, response) VALUES (?, ?)", payload, response)
	if err != nil {
		log.Printf("Error inserting into cache: %v", err)
	}
}
