package commentcontroller

import (
	"net/http"
	"strconv"
	"time"

	comment "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Comment"
	"github.com/labstack/echo/v4"
)

//Register ...コメント登録
func Register(c echo.Context) error {

	var com comment.Comment

	//フロントからの各データ取得
	userID := c.FormValue("user_id")
	emotionID := c.FormValue("emotion_id")
	comm := c.FormValue("comment")
	latitude := c.FormValue("latitude")
	longtitude := c.FormValue("longtitude")
	prefecture := c.FormValue("prefecture")

	//str->float64->float32
	var lati float64
	lati, _ = strconv.ParseFloat(latitude, 32)
	var a float32 = float32(lati)

	var long float64
	long, _ = strconv.ParseFloat(longtitude, 32)
	var b float32 = float32(long)

	//構造体に入れる
	com.UserID, _ = strconv.Atoi(userID)
	com.EmotionID, _ = strconv.Atoi(emotionID)
	com.Comment = comm
	com.Latitude = a
	com.Longtitude = b
	com.Prefecture = prefecture
	com.DateTime = time.Now()

	//DB処理
	hasSuccess := comment.Create(com)
	returnValue := map[string]bool{"hasSuccess": hasSuccess}
	return c.JSON(http.StatusOK, returnValue)
}

//Delete ...コメントの削除
func Delete(c echo.Context) error {

	commentID := c.FormValue("comment_id")
	comdel, _ := strconv.Atoi(commentID)

	//DB処理
	hasSuccess := comment.Delete(comdel)
	returnValue := map[string]bool{"hasSuccess": hasSuccess}

	return c.JSON(http.StatusOK, returnValue)
}
