package main

import (
	"github.com/MeguMan/url-shortener/internal/app/apiserver"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connStr := "user=postgres password=12345 dbname=restapi_test sslmode=disable"

	if err := apiserver.Start(connStr); err != nil {
		log.Fatal(err)
	}
}
