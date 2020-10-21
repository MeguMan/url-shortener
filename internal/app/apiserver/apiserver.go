package apiserver

import (
	"database/sql"
	"github.com/MeguMan/url-shortener/internal/app/store/postgres_store"
	"net/http"
)

//Start ...
func Start(DatabaseURL string) error {
	db, err := newDB(DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	s := postgres_store.New(db)
	server := NewServer(s)

	return http.ListenAndServe(":8080", server)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
