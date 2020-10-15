package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Link struct {
	InitialLink   string
	ShortenedLink string
}

func (l *Link) Validate() error {
	return validation.ValidateStruct(l, validation.Field(&l.InitialLink, validation.Required, is.URL))
}
