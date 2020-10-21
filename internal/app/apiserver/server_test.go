package apiserver_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MeguMan/url-shortener/internal/app/apiserver"
	"github.com/MeguMan/url-shortener/internal/app/model"
	"github.com/MeguMan/url-shortener/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_CreateLink(t *testing.T) {
	s := apiserver.NewServer(teststore.New())
	testCases := []struct {
		name         string
		link         model.Link
		expectedCode int
	}{
		{
			name: "valid",
			link: model.Link{
				InitialLink: "yandex.ru",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid link",
			link: model.Link{
				InitialLink: "google",
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestBody, err := json.Marshal(tc.link)
			if err != nil {
				panic(err)
			}
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/createlink", bytes.NewBuffer(requestBody))
			defer req.Body.Close()
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_Redirect(t *testing.T) {
	s := apiserver.NewServer(teststore.New())
	testCases := []struct {
		name         string
		link         model.Link
		expectedCode int
	}{
		{
			name: "valid",
			link: model.Link{
				ShortenedLink: "12345",
			},
			expectedCode: http.StatusTemporaryRedirect,
		},
		{
			name: "link not found",
			link: model.Link{
				ShortenedLink: "1234567",
			},
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", tc.link.ShortenedLink), nil)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
