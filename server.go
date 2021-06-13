package main

import (
	"go-learn-cicd/calculation"
	"go-learn-cicd/middleware"
	"go-learn-cicd/models"

	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// load env
	loadEnv()

	// load middleware
	e := echo.New()
	middleware.InitMiddleware(e)

	// load function package
	calcUC := calculation.InitAdd()

	e.GET("/hello", func(c echo.Context) error {
		hw := models.HelloWorld{Message: "Hello World!"}
		return c.JSON(http.StatusOK, hw)
	})

	e.POST("/add", func(c echo.Context) error {
		p := new(models.AddPayload)
		if err := c.Bind(p); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := c.Validate(p); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		resp := calcUC.Add(*p)
		return c.JSON(http.StatusOK, resp)
	})

	e.Logger.Fatal(e.Start(os.Getenv(`PORT`)))
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
