package http_server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Log the incoming request
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		// Log the completed request with duration
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If the path is not root and ends with a slash, remove it
		if r.URL.Path != "/" && strings.HasSuffix(r.URL.Path, "/") {
			http.Redirect(w, r, strings.TrimSuffix(r.URL.Path, "/"), http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}

type ControllerInterface interface {
	RegisterHandler(path string, handler func(http.ResponseWriter, *http.Request))
}

type Server struct {
	httpServer *http.Server
	mux        *http.ServeMux
}

func (s *Server) RegisterHandler(path string, handler func(http.ResponseWriter, *http.Request)) {
	s.mux.HandleFunc(path, handler)
}

func (s *Server) Listen() error {
	fmt.Println("Server is running on port 3000")
	err := s.httpServer.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}

func CreateServer(port string) *Server {
	mux := http.NewServeMux()
	wrappedMux := removeTrailingSlash(mux)
	loggedMux := loggerMiddleware(wrappedMux)
	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: loggedMux,
	}
	return &Server{
		httpServer: httpServer,
		mux:        mux,
	}
}
