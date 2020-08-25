package emotion

import db "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB"

//Emotion ...table: Emotionsのモデル
type Emotion struct {
	ID      int
	Emotion int
	Point   int
}

//Create ...Emotionモデルの保存
func Create(emotionNum int, point int) *Emotion {
	db := db.Connect()
	defer db.Close()

	emotion := Emotion{Emotion: emotionNum, Point: point}
	db.Create(&emotion)

	return &emotion
}

//Get ...Emotionモデルの取得
func Get(id int) *Emotion {
	db := db.Connect()
	defer db.Close()

	emotion := Emotion{}
	db.First(&emotion, id)

	return &emotion
}
