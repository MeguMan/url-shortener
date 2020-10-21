package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/MeguMan/url-shortener/internal/app/model"
	"github.com/MeguMan/url-shortener/internal/app/store"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"time"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/{shortenedLink}", s.redirect()).Methods("GET")
	s.router.HandleFunc("/createlink", s.createLink()).Methods("POST")
}

func (s *server) redirect() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		shortenedLink := vars["shortenedLink"]

		lr := s.store.Link()

		l, err := lr.FindByShortenedLink(shortenedLink)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			fmt.Println(err)
			return
		}
		l.AddProtocol()

		http.Redirect(w, r, l.InitialLink, http.StatusTemporaryRedirect)
	}
}

func (s *server) createLink() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		l := model.Link{}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewDecoder(r.Body).Decode(&l)

		err := l.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		l.ShortenedLink = RandString(7)
		lr := s.store.Link()

		err = lr.Create(&l)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(l.ShortenedLink)
	}
}

func RandString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678"

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
