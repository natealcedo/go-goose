package main

import (
	"github.com/natealcedo/go-goose/controllers"
	"github.com/natealcedo/go-goose/http-server"
	"github.com/natealcedo/go-goose/models"
	"github.com/natealcedo/go-goose/repository"
	"github.com/natealcedo/go-goose/services"
	"net/http"
)

func main() {
	db, err := models.CreateDatabaseClient()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	server := http_server.CreateServer("3000")

	postService := services.NewPostService(repository.NewGormRepository[models.Post](db.DB))
	postController := controllers.NewController(postService, server)

	// Register dynamic route
	postController.RegisterMethodHandlers("/posts/{id}", map[string]func(http.ResponseWriter, *http.Request){
		"GET":    postController.GetByID,
		"DELETE": postController.DeleteByID,
	}, true)

	// Register static route
	postController.RegisterMethodHandlers("/posts", map[string]func(http.ResponseWriter, *http.Request){
		"GET":  postController.Get,
		"POST": postController.POST,
	}, false)

	err = server.Listen()

	if err != nil {
		panic(err)
	}

}
