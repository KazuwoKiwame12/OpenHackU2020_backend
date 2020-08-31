package main

import (
	"net/http"
	"os"

	commentlistcontroller "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Controller/CommentListController"
	deletedata "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Controller/DeleteData"
	prefecturecontroller "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Controller/PrefectureController"
	responselistcontroller "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Controller/ResponseListController"
	usercontroller "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Controller/UserController"
	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")

	e := echo.New()
	e.GET("/", helloWorld)
	e.GET("/emotion/prefectures", prefecturecontroller.PrefectureInfoList)
	e.GET("/emotion/:prefecture/comments", commentlistcontroller.CommentsInPrefecture)
	e.GET("/emotion/:prefecture/comments/:comment_id", responselistcontroller.ResponseListInComment)
	e.POST("/user/register", usercontroller.Register)
        //e.POST("/comment/register", registerdata.CommentRegister)
        //e.POST("/comment/response", registerdata.ResponseRegister)
	e.DELETE("/comment/delete", deletedata.CommentDelete)
	e.DELETE("/comment/response/delete", deletedata.ResponseDelete)
	e.PATCH("/user/edit", usercontroller.Edit)
	e.Logger.Fatal(e.Start(":" + port))
}

func helloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World!!")
}
