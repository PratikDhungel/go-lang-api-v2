package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/quality-dashboard/config"
	"log"
)


func ConnectDB() (*sqlx.DB, error){

	//dbDriver := "mysql"
	//dbUser := "root"
	//dbPass := ""
	//dbName := "QualityDashboard"

	//DB, err = sql.Open(dbDriver, "root@tcp(127.0.0.1:3306)/QualityDashboard")

	db, err := sqlx.Connect(config.DBDRIVER, config.DBURL)

	fmt.Println("Connecting to DB")

	fmt.Println(db)

	if(err != nil){
		log.Fatal(err)
	}

	return db, err
}
