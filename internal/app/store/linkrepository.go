package store

import "github.com/MeguMan/url-shortener/internal/app/model"

type LinkRepository interface {
	Create(*model.Link) error
	FindByShortenedLink(string) (*model.Link, error)
}
