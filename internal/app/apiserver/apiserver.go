package apiserver

import (
	"database/sql"
	store2 "github.com/MeguMan/url-shortener/internal/app/store/sqlstore"
	"net/http"
)

//Start ...
func Start(DatabaseURL string) error {

	db, err := newDB(DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	s := store2.New(db)
	server := newServer(s)

	return http.ListenAndServe(":8181", server)
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
