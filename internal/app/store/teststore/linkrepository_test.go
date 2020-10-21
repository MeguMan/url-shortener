package teststore_test

import (
	"github.com/MeguMan/url-shortener/internal/app/model"
	"github.com/MeguMan/url-shortener/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkRepository_Create(t *testing.T) {
	s := teststore.New()
	l := model.TestLink(t)
	assert.NoError(t, s.Link().Create(l))
	assert.NotNil(t, l.ID)
}

func TestLinkRepository_FindByShortenedLink(t *testing.T) {
	s := teststore.New()
	tl := model.TestLink(t)
	l, err := s.Link().FindByShortenedLink(tl.ShortenedLink)
	assert.NoError(t, err)
	assert.NotNil(t, l)
}

//asd
