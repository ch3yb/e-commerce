package httpServer

import (
	"e-commerce/service"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type HttpServer struct {
	Addr    string
	Logger  *zap.Logger
	Service *service.Service
}

func Start(addr string, logger *zap.Logger, svc *service.Service) error {
	s := &HttpServer{
		Addr:    ":" + addr,
		Logger:  logger,
		Service: svc,
	}
	var router = s.NewRouter()
	var handler = cors(router)

	var server = http.Server{
		Addr:    s.Addr,
		Handler: handler,
	}
	log.Printf("Http Server (REST) starting on port %s", s.Addr)
	return server.ListenAndServe()
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", r.Method)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}
