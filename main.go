package main

import (
	"log"
	"net/http"
	"os"

	db "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB"
	"github.com/labstack/echo/v4"
)

func main() {

	port := os.Getenv("PORT")

	e := echo.New()
	e.GET("/", helloWorld)
	e.Logger.Fatal(e.Start(":" + port))
}

func helloWorld(c echo.Context) error {
	db := db.Connect()
	defer db.Close()
	log.Print(db)

	return c.JSON(http.StatusOK, "Hello World!!")
}
