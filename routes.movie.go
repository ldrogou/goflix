package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type JsonMovie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Duration    int    `json:"duration"`
	TrailerURL  string `json:"trailer_url"`
}

func (s *server) handleMovieList() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		movies, err := s.store.GetMovies()
		if err != nil {
			log.Printf("Cannot load Movies (err=%v)", err)
			s.response(rw, r, nil, http.StatusInternalServerError)

			return
		}

		var resp = make([]JsonMovie, len(movies))
		for i, m := range movies {
			resp[i] = mapMovieToJson(m)
		}

		s.response(rw, r, resp, http.StatusOK)

	}
}

func (s *server) handleMovie() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		vars, _ := mux.Vars(r)["id"]
		movieID, err := strconv.ParseInt(vars, 10, 64)
		if err != nil {
			log.Printf("Cannot parse id (err=%v)", err)
			s.response(rw, r, nil, http.StatusBadRequest)

			return
		}
		movie, err := s.store.GetMovieByID(movieID)
		if err != nil {
			log.Printf("Cannot load Movie (err=%v)", err)
			s.response(rw, r, nil, http.StatusNotFound)

			return
		}
		if movie == nil {
			s.response(rw, r, nil, http.StatusNotFound)

			return
		}

		var resp = mapMovieToJson(movie)
		s.response(rw, r, resp, http.StatusOK)
	}
}

func (s *server) handleInsertMovie() http.HandlerFunc {
	type request struct {
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		Duration    int    `json:"duration"`
		TrailerURL  string `json:"trailer_url"`
	}
	return func(rw http.ResponseWriter, r *http.Request) {

		req := request{}
		err := s.decode(rw, r, &req)
		if err != nil {
			log.Printf("Cannot parse movie body err=%v", err)
			s.response(rw, r, nil, http.StatusBadGateway)
			return
		}

		//Create movie
		m := &Movie{
			ID:          0,
			Title:       req.Title,
			ReleaseDate: req.ReleaseDate,
			Duration:    req.Duration,
			TrailerURL:  req.TrailerURL,
		}

		//store movie in database
		err = s.store.CreateMovie(m)
		if err != nil {
			log.Printf("Cannot create movie err=%v", err)
			s.response(rw, r, nil, http.StatusInternalServerError)
			return
		}
		var resp = mapMovieToJson(m)
		s.response(rw, r, resp, http.StatusOK)

	}
}

func mapMovieToJson(m *Movie) JsonMovie {
	return JsonMovie{
		ID:          m.ID,
		Title:       m.Title,
		ReleaseDate: m.ReleaseDate,
		Duration:    m.Duration,
		TrailerURL:  m.TrailerURL,
	}
}
