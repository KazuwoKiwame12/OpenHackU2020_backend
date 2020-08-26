package commentservice

import (
	"log"
	"time"

	comment "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Comment"
	user "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/User"
)

//CommentForClient ...クライアントに渡すデータ構造
type CommentForClient struct {
	ID         int
	Emotion    int
	Latitude   float32
	Longtitude float32
	UserName   string
	DateTime   time.Time
}

//ConvertCtoCFC ...[]commentを[]CommentForClientに変換する
func ConvertCtoCFC(commentList []comment.Comment) []CommentForClient {
	commentListForClient := []CommentForClient{}

	for index, comment := range commentList {
		log.Println(index)
		user := user.Get(comment.UserID)
		commentForClient := CommentForClient{}

		commentForClient.ID = comment.ID
		commentForClient.Emotion = comment.EmotionID
		commentForClient.Latitude = comment.Latitude
		commentForClient.Longtitude = comment.Longtitude
		commentForClient.UserName = user.Name
		commentForClient.DateTime = comment.DateTime

		commentListForClient = append(commentListForClient, commentForClient)
	}

	return commentListForClient
}
