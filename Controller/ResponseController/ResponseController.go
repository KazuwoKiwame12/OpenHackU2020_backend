package responsecontroller

import (
	"net/http"
	"strconv"
	"time"

	response "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Response"
	"github.com/labstack/echo/v4"
)

//Register ...コメントに対する返信登録
func Register(c echo.Context) error {

	//構造体宣言
	var res response.Response

	//フロントからのデータ取得
	userID := c.FormValue("user_id")
	commentID := c.FormValue("comment_id")
	comm := c.FormValue("comment")

	//構造体に入れる
	res.UserID, _ = strconv.Atoi(userID)
	res.CommentID, _ = strconv.Atoi(commentID)
	res.Comment = comm
	res.DateTime = time.Now()

	//DB処理
	hasSuccess := response.Create(res)
	returnValue := map[string]bool{"hasSuccess": hasSuccess}
	return c.JSON(http.StatusOK, returnValue)
}

//Delete ...返信の削除
func Delete(c echo.Context) error {

	responseID := c.FormValue("response_id")
	resdel, _ := strconv.Atoi(responseID)

	//DB処理
	hasSuccess := response.Delete(resdel)
	returnValue := map[string]bool{"hasSuccess": hasSuccess}

	return c.JSON(http.StatusOK, returnValue)
}
