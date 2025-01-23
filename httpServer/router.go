package httpServer

import "net/http"

func (s *HttpServer) NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /login", s.Login)
	return mux
}
