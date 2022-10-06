package conf

import (
	"bego/logger"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	instance *ServerConf
	confOnce sync.Once)

type ServerConf struct {
	dbhost string
	dbport string
	dbuser string
	dbpwd  string
	dbname   string	
	dbMaxIdleConns int
	dbMaxOpenConns int 
	dbConnMaxLifetime int	
	wPort string
	//count int
}


// Rectangle 를 반환하는 함수를 만들었다.
func CreateServerConf(envfilepath string) *ServerConf {
	confOnce.Do(func() {
		//&ServerConf{}		//new(ServerConf)
		instance = new(ServerConf)
		//instance := &ServerConf{}
		//instance.count = 0
		//instance.loadServerConf(envfilepath)		
		instance.loadServerConf(envfilepath)		
	})
		
	//db := DBInfo{}	
	return instance
	//return &db
}

// func (sconf *ServerConf) Add() {
// 	sconf.count++
// }

// func (sconf *ServerConf) Get() int {
// 	return sconf.count
// }

func (sconf *ServerConf) loadServerConf(envfilepath string) {

	err := godotenv.Load(envfilepath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	sconf.dbhost = os.Getenv("DBHost")
	sconf.dbport = os.Getenv("DBPort")
	sconf.dbuser = os.Getenv("DBUser")
	sconf.dbpwd = os.Getenv("DBPwd")
	sconf.dbname = os.Getenv("DBName")
	sconf.wPort = os.Getenv("WPort")
	
	sconf.dbMaxIdleConns, err = strconv.Atoi(os.Getenv("DBMaxIdleConns"))
	if err != nil {		
		logger.Error(err.Error())
		sconf.dbMaxIdleConns =2		
	}

	sconf.dbMaxOpenConns, err = strconv.Atoi(os.Getenv("DBMaxOpenConns"))
	if err != nil {
		logger.Error(err.Error())
		sconf.dbMaxOpenConns =2		
	}
	
	sconf.dbConnMaxLifetime, err = strconv.Atoi(os.Getenv("DBConnMaxLifetime"))
	if err != nil {
		logger.Error(err.Error())
		sconf.dbConnMaxLifetime = 5
	}

	log.Println(sconf.GetDBConnString())	
}

func (sconf *ServerConf) GetDBConnString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
	sconf.dbhost, sconf.dbuser, sconf.dbpwd, sconf.dbname, sconf.dbport)
}

func (sconf *ServerConf) GetWPort() string {
	return sconf.wPort
}

func (sconf *ServerConf) GetDBMaxIdleConns() int {
	return sconf.dbMaxIdleConns
}

func (sconf *ServerConf) GetDBMaxOpenConns() int {
	return sconf.dbMaxOpenConns
}

func (sconf *ServerConf) GetDBConnMaxLifetime() int {
	return sconf.dbConnMaxLifetime
}
	 