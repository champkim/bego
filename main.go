package main

import (
	"bego/app"
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

type DBInfo struct {
	Host string 
	Port string 
	User string 
	Pwd  string 
	DB   string 
}

func main() {
	dbInfo := &DBInfo{      // 값 설정
        Host: "127.0.0.1",
		Port: "5432",
		User: "ontune",
		Pwd: "ontune",
		DB: "ontune",
    }

	dbConn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
	dbInfo.Host, dbInfo.User, dbInfo.Pwd, dbInfo.DB, dbInfo.Port)

	//mux := app.MakeHandler("./test.db") //flag.Args 이런걸로 사용하자. 설정인자는 최대한 바깥으로 빼자
	//os.Getenv("DATABASE_URL")
	mux := app.MakeHandler(dbConn)
	defer mux.Close() //finally 개념

	ngri := negroni.Classic()
	ngri.UseHandler(mux)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", ngri)
	if err != nil {
		panic(err)
	}
}