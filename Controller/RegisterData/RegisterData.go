package registerdata

import (
	"net/http"
	"strconv"
	"time"

	comment "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Comment"
	"github.com/labstack/echo/v4"
)

//Register---初期コメント及び情報登録
func CommnetRegister(c echo.Context) error {

	var com comment.Comment

	//フロントからの各データ取得
	id := c.FormValue("user_id")
	emotion := c.FormValue("emotion")
	comm := c.FormValue("comment")
	latitude := c.FormValue("latitude")
	longtitude := c.FormValue("longtitude")
	prefecture := c.FormValue("prefecture")
	datatime := c.FormValue("dateTime")

	//str->float64->float32
	var lati float64
	lati, _ = strconv.ParseFloat(latitude, 32)
	var a float32 = float32(lati)

	var long float64
	long, _ = strconv.ParseFloat(longtitude, 32)
	var b float32 = float32(long)

	//time型の処理
	str := datatime
	layout := "Mon, 2 Jan 2006 15:04:05 -0700"
	t, _ := time.Parse(layout, str)

	//構造体に入れる
	com.ID, _ = strconv.Atoi(id)
	com.EmotionID, _ = strconv.Atoi(emotion)
	com.Comment = comm
	com.Latitude = a
	com.Longtitude = b
	com.Prefecture = prefecture
	com.DateTime = t
	comment.Create(com)

	return c.JSON(http.StatusOK, "GOOD")
}

func ResponseRegister(c echo.Context) error {

	return c.JSON(http.StatusOK, "GOOD")

}
