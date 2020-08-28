package prefectureservice

import (
	"image/color"
	"strconv"

	comment "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Comment"
	commentservice "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Service/CommentService"
)

//PrefectureInfo ...クライアントに返すデータ
type PrefectureInfo struct {
	Prefecture string
	Color      string
}

//GetPrefectureInfoList ...各県の色と名前の情報取得
func GetPrefectureInfoList() []PrefectureInfo {
	prefectureInfoList := []PrefectureInfo{}
	for _, prefecture := range getAllPrefectures() {
		commentList := comment.GetListByPrefecture(prefecture)
		emotionList := commentservice.GetCommentEmotionList(commentList)
		prefectureColor := calcColor(emotionList)

		prefectureInfo := PrefectureInfo{Prefecture: prefecture, Color: prefectureColor}
		prefectureInfoList = append(prefectureInfoList, prefectureInfo)
	}

	return prefectureInfoList
}

//GetAllPrefectures ...全ての県名を取得
func getAllPrefectures() [47]string {
	prefectureList := [47]string{
		"HOKKAIDO",
		"AOMORI",
		"IWATE",
		"MIYAGI",
		"AKITA",
		"YAMAGATA",
		"FUKUSHIMA",
		"IBARAKI",
		"TOCHIGI",
		"GUNMA",
		"SAITAMA",
		"CHIBA",
		"TOKYO",
		"KANAGAWA",
		"NIIGATA",
		"TOYAMA",
		"ISHIKAWA",
		"FUKUI",
		"YAMANASHI",
		"NAGANO",
		"GIFU",
		"SHIZUOKA",
		"AICHI",
		"MIE",
		"SHIGA",
		"KYOTO",
		"OSAKA",
		"HYOGO",
		"NARA",
		"WAKAYAMA",
		"TOTTORI",
		"SHIMANE",
		"OKAYAMA",
		"HIROSHIMA",
		"YAMAGUCHI",
		"TOKUSHIMA",
		"KAGAWA",
		"EHIME",
		"KOCHI",
		"OKINAWA",
		"FUKUOKA",
		"SAGA",
		"NAGASAKI",
		"KUMAMOTO",
		"OITA",
		"MIYAZAKI",
		"KAGOSHIMA",
	}

	return prefectureList
}

func calcColor(emotionList []int) string {
	colorList := [4]color.NRGBA{}
	colorList[0] = color.NRGBA{R: 255, G: 102, B: 255, A: 160} //喜
	colorList[1] = color.NRGBA{R: 255, G: 51, B: 51, A: 160}   //怒
	colorList[2] = color.NRGBA{R: 153, G: 153, B: 255, A: 160} //哀
	colorList[3] = color.NRGBA{R: 204, G: 204, B: 51, A: 160}  //楽

	mixColor := color.NRGBA{}

	for index, emotion := range emotionList {
		colorParam := convertEIDtoColorParam(emotion)
		if index == 0 {
			mixColor = colorList[colorParam]
		} else {
			alfaRate := (colorList[colorParam].A / 255)
			newR := colorList[colorParam].R*alfaRate + mixColor.R*(1-alfaRate)
			newG := colorList[colorParam].G*alfaRate + mixColor.G*(1-alfaRate)
			newB := colorList[colorParam].B*alfaRate + mixColor.B*(1-alfaRate)
			mixColor = color.NRGBA{R: newR, G: newG, B: newB, A: 205}
		}
	}

	stringR := strconv.FormatInt(int64(mixColor.R), 16)
	stringG := strconv.FormatInt(int64(mixColor.G), 16)
	stringB := strconv.FormatInt(int64(mixColor.B), 16)

	returnColor := stringR + stringG + stringB

	return returnColor
}

func convertEIDtoColorParam(emotionID int) int {
	result := 0
	if 1 <= emotionID && emotionID <= 3 {
		result = 0
	} else if 4 <= emotionID && emotionID <= 6 {
		result = 1
	} else if 7 <= emotionID && emotionID <= 9 {
		result = 2
	} else if 10 <= emotionID && emotionID <= 12 {
		result = 3
	}

	return result
}
