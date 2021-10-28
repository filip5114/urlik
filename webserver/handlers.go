package webserver

import "net/http"

func (s *server) mainHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Elo!"))
	}
}
