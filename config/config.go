package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var PORT int
var DBDRIVER, DBURL string

func LoadEnv(){

	var err error
	err = godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))

	if(err != nil){
		log.Fatal(err)
	}

	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("%s:@/%s", os.Getenv("DB_USER"), os.Getenv("DB_NAME"))
}



