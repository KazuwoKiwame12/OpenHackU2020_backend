package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Connect ...DBとの接続
func Connect() *gorm.DB {
	/*ローカルではこちらを用いる
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	connstr := dbURI
	*/
	connstr := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
