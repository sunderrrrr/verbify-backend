package main

import (
	"WhyAi"
	"WhyAi/pkg/handler"
	"WhyAi/pkg/repository"
	"WhyAi/pkg/service"
	"WhyAi/pkg/utils/logger"
	"fmt"
	"os"
)

func main() {
	/*if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}*/
	fmt.Println("Initializing...")
	fmt.Println("\n██╗   ██╗███████╗██████╗ ██████╗ ██╗███████╗██╗   ██╗\n██║   ██║██╔════╝██╔══██╗██╔══██╗██║██╔════╝╚██╗ ██╔╝\n██║   ██║█████╗  ██████╔╝██████╔╝██║█████╗   ╚████╔╝ \n╚██╗ ██╔╝██╔══╝  ██╔══██╗██╔══██╗██║██╔══╝    ╚██╔╝  \n ╚████╔╝ ███████╗██║  ██║██████╔╝██║██║        ██║   \n  ╚═══╝  ╚══════╝╚═╝  ╚═╝╚═════╝ ╚═╝╚═╝        ╚═╝   \n                                                     \n")
	fmt.Println("Version: 1.0.0")
	db, err := repository.NewDB(repository.DB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL"),
	})
	if err != nil {
		logger.Log.Fatalf("Error connecting to database: %v", err)
	}
	logger.Log.Println("Connecting to database")
	NewRepository := repository.NewRepository(db)
	NewService := service.NewService(NewRepository)
	NewHandler := handler.NewHandler(NewService)
	server := new(WhyAi.Server)
	logger.Log.Println("Running server")
	if err = server.Run(os.Getenv("SERVER_PORT"), NewHandler.InitRoutes(os.Getenv("FRONTEND_URL"))); err != nil {
		logger.Log.Fatalf("Fatal Error: %v", err)
	}
}
