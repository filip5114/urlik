package webserver

import (
	"fmt"
	"net/http"
)

type server struct {
	addr string
	mux  *http.ServeMux
}

func Run(port int) error {
	s := &server{
		addr: fmt.Sprintf(":%v", port),
		mux:  http.NewServeMux(),
	}
	s.routes()
	return http.ListenAndServe(s.addr, s)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
