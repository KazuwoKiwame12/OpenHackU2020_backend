package comment

import (
	"time"

	db "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB"
)

//Comment ...table: Commentsのモデル
type Comment struct {
	ID         int
	UserID     int
	EmotionID  int
	Prefecture string
	Latitude   float32
	Longtitude float32
	Comment    string
	DateTime   time.Time
}

//Create ...Commentsモデルの保存
func Create(comment Comment) *Comment {
	db := db.Connect()
	defer db.Close()

	db.Create(&comment)

	return &comment
}

//Get ...Commentsモデルの取得
func Get(id int) *Comment {
	db := db.Connect()
	defer db.Close()

	comment := Comment{}
	db.First(&comment, id)

	return &comment
}

//Delete ...Commentsモデルの削除
func Delete(id int) *Comment {
	db := db.Connect()
	defer db.Close()

	comment := Comment{}
	comment.ID = id
	db.Delete(&comment)

	return &comment
}
