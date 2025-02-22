package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client
var db *sql.DB

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	Addr, _ := os.LookupEnv("REDIS_ADDR")
	Pass, _ := os.LookupEnv("REDIS_PASSWORD")
	rdb = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Pass,
		DB:       0,
	})
	connectSQL()
	setUpSQLITE()
	InitCache()
}

var QUEUE_NAME = "queue"

func producer(wg *sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter data to push (or 'exit' to quit): ")
		if !scanner.Scan() {
			break
		}
		data := scanner.Text()
		if strings.ToLower(data) == "exit" {
			break
		}
		rdb.RPush(ctx, QUEUE_NAME, data)
		fmt.Println("Pushed:", data)
	}
}
func producerFromFile(wg *sync.WaitGroup, filename string) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.TrimSpace(scanner.Text()) // Remove leading/trailing whitespace
		if data == "" {
			continue // Skip empty lines
		}

		rdb.RPush(ctx, QUEUE_NAME, data)
		fmt.Println("Pushed:", data)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

// ProcessCSV reads a CSV file, extracts the first column, and pushes it to Redis queue
func consumer(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		link, err := rdb.LPop(ctx, QUEUE_NAME).Result()
		if err == redis.Nil {
			continue
		} else if err != nil {
			fmt.Println("Error fetching from queue:", err)
			return
		}
		index(link)
	}

}
func run() {
	var wg sync.WaitGroup

	wg.Add(1)
	//go producer(&wg)
	go consumer(&wg)

	wg.Wait()
}
func main() {
	//ClearQueue()
	// var wg sync.WaitGroup
	// wg.Add(1)
	// //go producerFromFile(&wg, "data.txt")
	// go consumer(&wg)
	// wg.Wait()
	app := setupServer()
	log.Fatal(app.Listen(":3000"))
	fmt.Println(bm25("Chess in react"))
}
