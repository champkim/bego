package app

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBInfo struct {
	Host string
	Port string
	User string
	Pwd  string
	DB   string
	WPort string
}

// Rectangle 를 반환하는 함수를 만들었다.
func CreateDBInfo(envfilepath string) *DBInfo {
	//db := DBInfo{}
	db := new(DBInfo)
	db.loadDBConf(envfilepath)
	return db
	//return &db
}

func (db *DBInfo) loadDBConf(envfilepath string) {

	err := godotenv.Load(envfilepath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	db.Host = os.Getenv("Host")
	db.Port = os.Getenv("Port")
	db.User = os.Getenv("User")
	db.Pwd = os.Getenv("Pwd")
	db.DB = os.Getenv("DB")
	db.WPort = os.Getenv("WPort")

	log.Println(db.GetDBConnString())	
}

func (db *DBInfo) GetDBConnString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		db.Host, db.User, db.Pwd, db.DB, db.Port)
}