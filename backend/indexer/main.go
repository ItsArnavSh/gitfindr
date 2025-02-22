package main

import (
	"context"
	"database/sql"
	"log"
	"os"

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
func main() {

}
