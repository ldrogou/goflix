package main

import (
	"log"
	"net/http"
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

func mapMovieToJson(m *Movie) JsonMovie {
	return JsonMovie{
		ID:          m.ID,
		Title:       m.Title,
		ReleaseDate: m.ReleaseDate,
		Duration:    m.Duration,
		TrailerURL:  m.TrailerURL,
	}
}
