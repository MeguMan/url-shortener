package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"strings"
)

type Link struct {
	ID            int
	InitialLink   string
	ShortenedLink string
}

func (l *Link) Validate() error {
	return validation.ValidateStruct(l, validation.Field(&l.InitialLink, validation.Required, is.URL))
}

func (l *Link) AddProtocol() {
	if strings.Contains(l.InitialLink, "https://") || strings.Contains(l.InitialLink, "http://") {
		return
	}
	l.InitialLink = "https://" + l.InitialLink
}
