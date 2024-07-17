package controllers

import (
	"encoding/json"
	"github.com/natealcedo/go-goose/http-server"
	"github.com/natealcedo/go-goose/services"
	"net/http"
)

type Controller struct {
	service services.GenericService
	server  *http_server.Server
}

func NewController(service services.GenericService, server *http_server.Server) *Controller {
	return &Controller{
		service: service,
		server:  server,
	}
}

func (c *Controller) RegisterHandler(path string, handler func(http.ResponseWriter, *http.Request)) {
	c.server.RegisterHandler(path, handler)
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	items, err := c.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (c *Controller) POST(w http.ResponseWriter, r *http.Request) {
	var body interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.service.Create(body); err != nil {
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Successfully processed request"))
}
