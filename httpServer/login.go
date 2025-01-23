package httpServer

import (
	"log"
	"net/http"
)

func (s *HttpServer) Login(writer http.ResponseWriter, request *http.Request) {
	log.Println("Login RES: ", s.Service.Ping())
}
