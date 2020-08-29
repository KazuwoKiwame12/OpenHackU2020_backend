package deletedata

import (
	"net/http"
	"strconv"

	comment "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Comment"
	response "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Response"
	"github.com/labstack/echo/v4"
)

func CommentDelete(c echo.Context) error {

	//フロントからのデータ取得
	comment_id := c.FormValue("comment_id")
	comdel, _ := strconv.Atoi(comment_id)

	//DB処理
	hasSuccess := comment.Delete(comdel)
	returnValue := map[string]bool{"hasSuccess": hasSuccess}

	return c.JSON(http.StatusOK, returnValue)
}

func ResponseDelete(c echo.Context) error {

	//フロントからのデータ取得
	response_id := c.FormValue("response_id")
	resdel, _ := strconv.Atoi(response_id)

	//DB処理
	hasSuccess := response.Delete(resdel)
	returnValue := map[string]bool{"hasSuccess": hasSuccess}

	return c.JSON(http.StatusOK, returnValue)

}
