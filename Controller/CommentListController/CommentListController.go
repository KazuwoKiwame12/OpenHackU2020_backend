package commentlistcontroller

import (
	"net/http"

	comment "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Comment"
	commentservice "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Service/CommentService"
	"github.com/labstack/echo/v4"
)

//CommentListInPrefecture ...県に存在するコメント一覧取得
func CommentListInPrefecture(c echo.Context) error {
	commentList := comment.GetListByPrefecture(c.Param("prefecture"))
	commentListForClient := commentservice.ConvertCtoCFC(commentList)

	return c.JSON(http.StatusOK, commentListForClient)
}
