package usercontroller

import (
	"net/http"
	"strconv"

	user "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/User"
	"github.com/labstack/echo/v4"
)

/*Errorが返ってくるようになっていない*/
//Register ...ユーザ情報登録
func Register(c echo.Context) error {
	name := c.FormValue("name")
	user := user.Create(name)

	return c.JSON(http.StatusOK, user)
}

//Edit ...ユーザ情報編集
func Edit(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	newName := c.FormValue("newName")
	user := user.Update(id, newName)

	return c.JSON(http.StatusOK, user)
}
