package main

import (
	"bego/app"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)



func main() {
	

	dbInfo := CreateDBInfo()
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