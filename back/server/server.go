package server

import (
	"net/http"
	"github.com/gorilla/mux"
)

type httpServer struct {
	handlers Handlers
}

func NewHTTPServer(handlers Handlers) *httpServer {
	return &httpServer{
		handlers: handlers,
	}
}

func (s *httpServer) Run() error {
	router := mux.NewRouter()
	router.Use(corsMiddleware)

	router.Path("/analyse").Methods(http.MethodPost).HandlerFunc(s.handlers.HandleAnalyse)

	router.Path("/analyse").Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	fs := http.FileServer(http.Dir("./dist"))
  router.PathPrefix("/").Handler(fs)

	server := http.Server{
		Addr: ":3000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}