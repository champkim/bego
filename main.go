package main

import (
	"bego/app"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)



func main() {
	

	//dbInfo := app.CreateDBInfo()
	//mux := app.MakeHandler("./test.db") //flag.Args 이런걸로 사용하자. 설정인자는 최대한 바깥으로 빼자
	//os.Getenv("DATABASE_URL")	
	env := app.CreateDBInfo("server.env")
	mux := app.MakeHandler(env.GetDBConnString())
	defer mux.Close() //finally 개념

	ngri := negroni.Classic()
	ngri.UseHandler(mux)	

	log.Println("Started App")
	err := http.ListenAndServe(":" + env.GetWPort(), ngri)
	if err != nil {
		panic(err)
	}
}