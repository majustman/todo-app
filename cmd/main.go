package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

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

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
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
