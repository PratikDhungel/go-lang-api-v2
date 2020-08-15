package api

import (
	"fmt"
	"github.com/quality-dashboard/api/database"
	"github.com/quality-dashboard/api/router"
	"github.com/quality-dashboard/config"
	"log"
	"net/http"
)

func RunServer()  {

	config.LoadEnv()
	database.ConnectDB()
	fmt.Println("Running Server at port [::]", config.PORT)
	Listen(config.PORT)
}

func Listen(port int) {

	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
