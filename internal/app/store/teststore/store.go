package teststore

import (
	"github.com/MeguMan/url-shortener/internal/app/model"
	"github.com/MeguMan/url-shortener/internal/app/store"
)

type Store struct {
	linkRepository *LinkRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) Link() store.LinkRepository {
	if s.linkRepository != nil {
		return s.linkRepository
	}

	s.linkRepository = &LinkRepository{
		store: s,
		links: make(map[int]*model.Link),
	}

	return s.linkRepository
}
