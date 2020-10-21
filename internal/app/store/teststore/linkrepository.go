package teststore

import (
	"github.com/MeguMan/url-shortener/internal/app/model"
	"github.com/MeguMan/url-shortener/internal/app/store"
)

type LinkRepository struct {
	store *Store
	links map[int]*model.Link
}

func (lr *LinkRepository) Create(l *model.Link) error {
	if err := l.Validate(); err != nil {
		return err
	}

	lr.links = make(map[int]*model.Link)
	lr.links[l.ID] = l
	return nil
}

func (lr *LinkRepository) FindByShortenedLink(shortenedLink string) (*model.Link, error) {
	l := model.Link{
		InitialLink:   "google.com",
		ShortenedLink: "12345",
	}

	if l.ShortenedLink == shortenedLink {
		return &l, nil
	}

	return nil, store.ErrRecordNotFound
}
