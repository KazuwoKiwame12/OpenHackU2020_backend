package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	port := os.Getenv("PORT")

	e := echo.New()
	e.GET("/", helloWorld)
	e.Logger.Fatal(e.Start(":" + port))
}

func helloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World!!")
}
