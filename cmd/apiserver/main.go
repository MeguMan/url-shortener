package main

import (
	"github.com/MeguMan/go-http-api/internal/app/apiserver"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connStr := "user=postgres password=12345 dbname=restapi_test sslmode=disable"

	if err := apiserver.Start(connStr); err != nil {
		log.Fatal(err)
	}
}
