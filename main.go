package main

import (
	"Dapp_2/pkg/database"
	"Dapp_2/pkg/handler"
	"Dapp_2/pkg/handler/user"
	"Dapp_2/pkg/service/user"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("APP_ENV")
	db := database.New()

	server := handler.NewServer()

	userService := userservice.NewService(db, env)
	userhandler.NewHandler(server, "/user", userService)

}
