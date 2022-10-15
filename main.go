package main

import (
	"fmt"

	"github.com/dzgntutar/ccgin/pkg/config"
	"github.com/dzgntutar/ccgin/pkg/database"
	"github.com/dzgntutar/ccgin/pkg/router"
)

func main() {

	if err := config.Init(); err != nil {
		//config error
		fmt.Println("Config Error --> ", err)
	}
	conf := config.GetGlobalConfig()

	if err := database.Init(); err != nil {
		//db error
		fmt.Println("Db Error -->", err)
	}
	db := database.DB

	r := router.Init(db)
	r.Run(conf.Web.ServerName + ":" + conf.Web.Port)
}
