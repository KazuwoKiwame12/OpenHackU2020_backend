package registerdatacontroller

import (
	"net/http"
	"strconv"
	"time"

	comment "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Comment"
	response "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Response"
	"github.com/labstack/echo/v4"
)

//RegisterComment ...コメント登録
func RegisterComment(c echo.Context) error {

	var com comment.Comment

	//フロントからの各データ取得
	userID := c.FormValue("user_id")
	emotionID := c.FormValue("emotion_id")
	comm := c.FormValue("comment")
	latitude := c.FormValue("latitude")
	longtitude := c.FormValue("longtitude")
	prefecture := c.FormValue("prefecture")
	datetime := c.FormValue("dateTime")

	//str->float64->float32
	var lati float64
	lati, _ = strconv.ParseFloat(latitude, 32)
	var a float32 = float32(lati)

	var long float64
	long, _ = strconv.ParseFloat(longtitude, 32)
	var b float32 = float32(long)

	//time型の処理
	str := datetime
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, str)

	//構造体に入れる
	com.UserID, _ = strconv.Atoi(userID)
	com.EmotionID, _ = strconv.Atoi(emotionID)
	com.Comment = comm
	com.Latitude = a
	com.Longtitude = b
	com.Prefecture = prefecture
	com.DateTime = t

	//DB処理
	hasSuccess := comment.Create(com)
	returnValue := map[string]bool{"hasSuccess": hasSuccess}
	return c.JSON(http.StatusOK, returnValue)
}

//RegisterResponse ...コメントに対する返信登録
func RegisterResponse(c echo.Context) error {

	//構造体宣言
	var res response.Response

	//フロントからのデータ取得
	userID := c.FormValue("user_id")
	commentID := c.FormValue("comment_id")
	comm := c.FormValue("comment")
	datetime := c.FormValue("dateTime")

	//time型の処理
	str := datetime
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, str)

	//構造体に入れる
	res.UserID, _ = strconv.Atoi(userID)
	res.CommentID, _ = strconv.Atoi(commentID)
	res.Comment = comm
	res.DateTime = t

	//DB処理
	hasSuccess := response.Create(res)
	returnValue := map[string]bool{"hasSuccess": hasSuccess}
	return c.JSON(http.StatusOK, returnValue)
}
