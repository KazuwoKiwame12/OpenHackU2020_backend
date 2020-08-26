package main

import (
	"log"
	"net/http"
	"os"

	commentlistcontroller "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Controller/CommentListController"
	usercontroller "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Controller/UserController"
	db "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB"
	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")

	e := echo.New()
	e.GET("/", helloWorld)
	e.GET("/emotion/:prefecture/comments", commentlistcontroller.CommentsInPrefecture)
	e.POST("/user/register", usercontroller.Register)
	e.POST("/user/edit", usercontroller.Edit)
	e.Logger.Fatal(e.Start(":" + port))
}

func helloWorld(c echo.Context) error {
	db := db.Connect()
	defer db.Close()
	log.Print(db)

	return c.JSON(http.StatusOK, "Hello World!!")
}
