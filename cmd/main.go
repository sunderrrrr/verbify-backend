package main

import (
	"WhyAi"
	"WhyAi/pkg/handler"
	"WhyAi/pkg/repository"
	"WhyAi/pkg/service"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	//TODO Подвязать env
	db, err := repository.NewDB(repository.DB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL"),
	})
	if err != nil {
		log.Fatal("error connecting to database: %s", err.Error())
	}
	NewRepository := repository.NewRepository(db)
	NewService := service.NewService(NewRepository)
	NewHandler := handler.NewHandler(NewService)
	server := new(WhyAi.Server)
	err = server.Run(os.Getenv("SERVER_PORT"), NewHandler.InitRoutes())
	if err != nil {
		log.Fatal(err)
	}
}
