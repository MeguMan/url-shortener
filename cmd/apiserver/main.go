package main

import (
	"github.com/MeguMan/url-shortener/internal/app/apiserver"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	databaseUrl, exists := os.LookupEnv("DATABASE_URL")

	if exists {
		if err := apiserver.Start(databaseUrl); err != nil {
			log.Fatal(err)
		}
	}
}
