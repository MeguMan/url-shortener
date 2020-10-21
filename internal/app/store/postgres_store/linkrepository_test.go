package postgres_store_test

import (
	"github.com/MeguMan/url-shortener/internal/app/model"
	"github.com/MeguMan/url-shortener/internal/app/store"
	"github.com/MeguMan/url-shortener/internal/app/store/postgres_store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkRepository_Create(t *testing.T) {
	db, teardown := postgres_store.TestDB(t, "user=postgres password=12345 dbname=restapi_test sslmode=disable")
	defer teardown("links")
	s := postgres_store.New(db)
	l := model.TestLink(t)
	assert.NoError(t, s.Link().Create(l))
	assert.NotNil(t, l.ID)
}

func TestLinkRepository_FindByShortenedLink(t *testing.T) {
	db, teardown := postgres_store.TestDB(t, "user=postgres password=12345 dbname=restapi_test sslmode=disable")
	defer teardown("links")

	s := postgres_store.New(db)
	l1 := model.TestLink(t)
	_, err := s.Link().FindByShortenedLink(l1.ShortenedLink)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.Link().Create(l1)
	l2, err := s.Link().FindByShortenedLink(l1.ShortenedLink)
	assert.NoError(t, err)
	assert.NotNil(t, l2)
}
