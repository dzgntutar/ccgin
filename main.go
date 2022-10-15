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

	// fmt.Println("Web -->", conf.Web)
	// fmt.Println("DB -->", conf.Database)
	// fmt.Printf("%T", conf)
	// fmt.Println()
	// fmt.Printf(" %#v", conf)

	if err := database.Init(); err != nil {
		//db error
		fmt.Println("Db Error -->", err)
	}
	db := database.DB

	// fmt.Println("main-Db -->", db)
	// fmt.Println()

	r := router.Init(db)
	r.Run(conf.Web.ServerName + ":" + conf.Web.Port)
}
