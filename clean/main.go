package main

import (
	"go-echo/internal/middleware"
	"go-echo/internal/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	middleware.Init(e)
	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
