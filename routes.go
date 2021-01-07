package main

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/api/movies/{id:[0-9]+}", s.handleMovie()).Methods("GET")
	s.router.HandleFunc("/api/movies", s.handleMovieList()).Methods("GET")
	s.router.HandleFunc("/api/movies", s.handleInsertMovie()).Methods("POST")

}
