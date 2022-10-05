package conf

import "log"

func WriteDBErr(err error) {
	//panic(err)
	log.Fatal(err)
}