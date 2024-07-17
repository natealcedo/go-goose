package http_server

import (
	"fmt"
	"net/http"
)

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
	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	return &Server{
		httpServer: httpServer,
		mux:        mux,
	}
}
