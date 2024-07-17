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

	testTableService := services.NewTestTableService(repository.NewGormRepository[models.TestTable](db.DB))
	testController := controllers.NewController(testTableService, server)
	testController.RegisterMethodHandlers("/test", map[string]func(http.ResponseWriter, *http.Request){
		"GET":  testController.Get,
		"POST": testController.POST,
	})

	err = server.Listen()

	if err != nil {
		panic(err)
	}

}
