package registerdata

import (
	"net/http"
	"strconv"
	"time"

	comment "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Comment"
	response "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Response"
	"github.com/labstack/echo/v4"
)

//Register---初期コメント及び情報登録
func CommentResister(c echo.Context) error {

	var com comment.Comment

	//フロントからの各データ取得
	user_id := c.FormValue("user_id")
	emotion := c.FormValue("emotion")
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
	com.ID, _ = strconv.Atoi(user_id)
	com.EmotionID, _ = strconv.Atoi(emotion)
	com.Comment = comm
	com.Latitude = a
	com.Longtitude = b
	com.Prefecture = prefecture
	com.DateTime = t

	//DB処理
	comment.Create(com)

	return c.JSON(http.StatusOK, "true")
}

func ResponseRegister(c echo.Context) error {

	//構造体宣言
	var res response.Response

	//フロントからのデータ取得
	user_id := c.FormValue("user_id")
	comment_id := c.FormValue("comment_id")
	comm := c.FormValue("comment")
	datetime := c.FormValue("dateTime")

	//time型の処理
	str := datetime
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, str)

	//構造体に入れる
	res.UserID, _ = strconv.Atoi(user_id)
	res.CommentID, _ = strconv.Atoi(comment_id)
	res.Comment = comm
	res.DateTime = t

	//DB処理
	response.Create(res)

	return c.JSON(http.StatusOK, "true")

}
