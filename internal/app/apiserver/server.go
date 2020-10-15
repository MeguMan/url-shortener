package apiserver

import (
	"encoding/json"
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

func newServer(store *store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  *store,
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
		db := s.store.Db
		vars := mux.Vars(r)
		shortenedLink := vars["shortenedLink"]

		rows, err := db.Query("select * from links where shortened_link = $1", shortenedLink)
		if err != nil {
			panic(err)
		}

		l := model.Link{}
		for rows.Next() {
			err = rows.Scan(&l.InitialLink, &l.ShortenedLink)
			if err != nil {
				panic(err)
			}
		}

		http.Redirect(w, r, l.InitialLink, http.StatusSeeOther)
	}
}

func (s *server) createLink() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		db := s.store.Db
		link := model.Link{}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewDecoder(r.Body).Decode(&link)
		err := link.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		link.ShortenedLink = RandString(7)

		_, err = db.Exec("insert into links (initial_link, shortened_link) values ($1, $2)",
			link.InitialLink, link.ShortenedLink)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(link.ShortenedLink)
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
