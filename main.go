package main

import (
	"bego/app"
	"bego/conf"
	"bego/logger"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)



func main() {	

	//dbInfo := app.CreateDBInfo()
	//mux := app.MakeHandler("./test.db") //flag.Args 이런걸로 사용하자. 설정인자는 최대한 바깥으로 빼자
	//os.Getenv("DATABASE_URL")	
	sconf := conf.CreateServerConf("server.env")
	mux := app.MakeHandler(sconf)
	defer mux.Close() //finally 개념

	ngri := negroni.Classic()
	ngri.UseHandler(mux)	

	//log.Println("Started App")
	logger.Info("Started App")
	err := http.ListenAndServe(":" + sconf.GetWPort(), ngri)
	if err != nil {				
		log.Fatal(err)
		panic(err)
	}
}