package main

import (
	"log"

	"github.com/majustman/todo-app/pkg/repository"
	"github.com/majustman/todo-app/pkg/service"

	"github.com/majustman/todo-app"
	"github.com/majustman/todo-app/pkg/handler"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
