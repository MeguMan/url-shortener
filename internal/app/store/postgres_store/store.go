package postgres_store

import (
	"database/sql"
	"github.com/MeguMan/url-shortener/internal/app/store"
	_ "github.com/lib/pq" // ...
)

type Store struct {
	db             *sql.DB
	linkRepository *LinkRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Link() store.LinkRepository {
	if s.linkRepository != nil {
		return s.linkRepository
	}

	s.linkRepository = &LinkRepository{
		store: s,
	}

	return s.linkRepository
}
