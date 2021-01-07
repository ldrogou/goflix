package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	store  Store
}

func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *server) serveHTTP(rw http.ResponseWriter, r *http.Request) {
	logRequestMiddleware(s.router.ServeHTTP).ServeHTTP(rw, r)
}

func (s *server) response(rw http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	rw.Header().Add("Content-type", "application/json")
	rw.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		log.Printf("Cannot encode to json (err=%v)\n", err)
	}

}

func (s *server) decode(rw http.ResponseWriter, r *http.Request) interface{} {
	var data interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return err
	}

	return data
}
