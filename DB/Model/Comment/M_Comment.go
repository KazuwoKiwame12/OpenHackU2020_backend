package comment

import (
	"time"

	db "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB"
	dateservice "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Service/DateService"
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
func Create(comment Comment) bool {
	db := db.Connect()
	defer db.Close()

	result := db.Create(&comment)
	if result.Error != nil {
		return false
	}
	return true
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
func Delete(id int) bool {
	db := db.Connect()
	defer db.Close()

	comment := Comment{}
	comment.ID = id
	result := db.Delete(&comment)
	if result.Error != nil {
		return false
	}
	return true
}

//GetByIDAndPrefecture ...指定した県に存在する指定してIDのコメントを返す
func GetByIDAndPrefecture(prefecture string, commentID int) Comment {
	db := db.Connect()
	defer db.Close()

	now := time.Now()
	from := dateservice.StartOfDay(now)
	to := dateservice.StartOfNextDay(now)

	comment := Comment{}
	db.Where("prefecture = ? AND id = ? AND date_time >= ? AND date_time < ?", prefecture, commentID, from, to).First(&comment)
	return comment
}

//GetListByPrefecture ...指定した県のコメント一覧
func GetListByPrefecture(prefecture string) []Comment {
	db := db.Connect()
	defer db.Close()

	now := time.Now()
	from := dateservice.StartOfDay(now)
	to := dateservice.StartOfNextDay(now)

	comments := []Comment{}
	db.Where("prefecture = ? AND date_time >= ? AND date_time < ?", prefecture, from, to).Find(&comments)
	return comments
}
