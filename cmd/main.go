package main

import (
	"log"

	"github.com/spf13/viper"

	"github.com/majustman/todo-app/pkg/repository"
	"github.com/majustman/todo-app/pkg/service"

	"github.com/majustman/todo-app"
	"github.com/majustman/todo-app/pkg/handler"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing confogs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
