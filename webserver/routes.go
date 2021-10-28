package webserver

func (s *server) routes() {
	s.mux.HandleFunc("/", s.mainHandler())
}
