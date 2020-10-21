package model_test

import (
	"github.com/MeguMan/url-shortener/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLink_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		l       func() *model.Link
		isValid bool
	}{
		{
			name: "valid",
			l: func() *model.Link {
				return model.TestLink(t)
			},
			isValid: true,
		},
		{
			name: "invalid link",
			l: func() *model.Link {
				l := model.TestLink(t)
				l.InitialLink = "google"
				return l
			},
			isValid: false,
		},
		{
			name: "empty link",
			l: func() *model.Link {
				l := model.TestLink(t)
				l.InitialLink = ""
				return l
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.l().Validate())
			} else {
				assert.Error(t, tc.l().Validate())
			}
		})
	}
}

func TestLink_CheckProtocol(t *testing.T) {
	testCases := []struct {
		l        func() *model.Link
		expected string
	}{
		{
			l: func() *model.Link {
				l := model.TestLink(t)
				l.AddProtocol()
				return l
			},
			expected: "https://google.com",
		},
	}

	for _, tc := range testCases {
		if tc.l().InitialLink != tc.expected {
			t.Error("Test failed: expected, {} received: {}", tc.l().InitialLink, tc.expected)
		}
	}
}
