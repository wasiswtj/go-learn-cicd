package main

import (
	"go-learn-cicd/models"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	hw := models.HelloWorld{Message: "Hello World!"}

	// load env
	loadEnv()

	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, hw)
	})
	e.Logger.Fatal(e.Start(os.Getenv(`PORT`)))
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
