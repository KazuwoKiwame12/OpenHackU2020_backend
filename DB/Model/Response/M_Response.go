package response

import (
	"time"

	db "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB"
)

//Response ...table: Responsesのモデル
type Response struct {
	ID        int
	UserID    int
	CommentID int
	Comment   string
	DateTime  time.Time
}

//Create ...Responsesモデルの保存
func Create(response Response) *Response {
	db := db.Connect()
	defer db.Close()

	db.Create(&response)

	return &response
}

//Get ...Responsesモデルの取得
func Get(id int) *Response {
	db := db.Connect()
	defer db.Close()

	response := Response{}
	db.First(&response, id)

	return &response
}

//Delete ...Responsesモデルの削除
func Delete(id int) *Response {
	db := db.Connect()
	defer db.Close()

	response := Response{}
	response.ID = id
	db.Delete(&response)

	return &response
}
