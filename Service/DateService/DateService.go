package dateservice

import (
	"log"
	"time"
)

//StartOfDay ...引数の開始日時を返す
func StartOfDay(t time.Time) time.Time {
	startOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	log.Print(startOfDay)
	return startOfDay
}

//StartOfNextDay ...引数の次の日の開始日時を返す
func StartOfNextDay(t time.Time) time.Time {
	startOfDay := StartOfDay(t)
	startOfNextDay := startOfDay.AddDate(0, 0, 1)
	log.Print(startOfNextDay)
	return startOfNextDay
}
