package model

import "testing"

func TestLink(t *testing.T) *Link {
	t.Helper()

	return &Link{
		InitialLink:   "google.com",
		ShortenedLink: "12345",
	}
}
