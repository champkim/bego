package app

import "fmt"

type DBInfo struct {
	Host string
	Port string
	User string
	Pwd  string
	DB   string
}

// Rectangle 를 반환하는 함수를 만들었다.
func CreateDBInfo() *DBInfo {
	//db := DBInfo{}
	db := new(DBInfo)
	db.LoadDBConf()
	return db
	//return &db
}

func (db *DBInfo) LoadDBConf() {
	db.Host = "127.0.0.1"
	db.Port = "5432"
	db.User = "ontune"
	db.Pwd = "ontune"
	db.DB = "ontune"
	// dbInfo := &DBInfo{      // 값 설정
	//     Host: "127.0.0.1",
	// 	Port: "5432",
	// 	User: "ontune",
	// 	Pwd: "ontune",
	// 	DB: "ontune",
	// }
}

func (db *DBInfo) GetDBConnString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		db.Host, db.User, db.Pwd, db.DB, db.Port)
}