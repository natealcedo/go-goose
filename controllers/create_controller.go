package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/natealcedo/go-goose/http-server"
	"github.com/natealcedo/go-goose/responses"
	"github.com/natealcedo/go-goose/services"
	"net/http"
	"regexp"
	"strings"
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

func (c *Controller) RegisterMethodHandlers(path string, handlers map[string]func(http.ResponseWriter, *http.Request), dynamic bool) {
	c.server.RegisterHandler(path, func(w http.ResponseWriter, r *http.Request) {
		if dynamic {
			// Use a regular expression to extract the ID from the path
			re := regexp.MustCompile(`[^/]+$`)
			pathID := re.FindString(r.URL.Path)

			if pathID != "" && handlers[r.Method] != nil {
				handlers[r.Method](w, r)
				return
			}
			http.NotFound(w, r)
		} else {
			if handler, exists := handlers[r.Method]; exists {
				handler(w, r)
				return
			}
			http.NotFound(w, r)
		}
	})
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
	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	entity, err := c.service.Create(body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path
	pathSegments := strings.Split(r.URL.Path, "/")
	// Ensure there is an ID part in the URL path
	if len(pathSegments) < 3 {
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}
	id := pathSegments[len(pathSegments)-1] // Assumes the ID is the last segment

	// Retrieve the item by ID using the service
	item, err := c.service.GetByID(id)

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(responses.NOT_FOUND)
		return
	}

	// Marshal the item into JSON and write the response
	jsonResponse, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (c *Controller) DeleteByID(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path
	pathSegments := strings.Split(r.URL.Path, "/")
	// Ensure there is an ID part in the URL path
	if len(pathSegments) < 3 {
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}
	id := pathSegments[len(pathSegments)-1] // Assumes the ID is the last segment

	// Retrieve the item by ID using the service
	err := c.service.DeleteByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(responses.NOT_FOUND)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
