package postgres_store

import (
	"database/sql"
	"github.com/MeguMan/url-shortener/internal/app/model"
	"github.com/MeguMan/url-shortener/internal/app/store"
)

type LinkRepository struct {
	store *Store
}

func (lr *LinkRepository) Create(l *model.Link) error {
	_, err := lr.store.db.Exec("insert into links (initial_link, shortened_link) values ($1, $2)",
		l.InitialLink, l.ShortenedLink)

	return err
}

func (lr *LinkRepository) FindByShortenedLink(shortenedLink string) (*model.Link, error) {
	l := &model.Link{}
	if err := lr.store.db.QueryRow(
		"SELECT * FROM links WHERE shortened_link = $1",
		shortenedLink,
	).Scan(
		&l.ID,
		&l.InitialLink,
		&l.ShortenedLink,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return l, nil
}
