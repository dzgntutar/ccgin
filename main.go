package main

import (
	"fmt"

	"github.com/dzgntutar/ccgin/pkg/config"
	"github.com/dzgntutar/ccgin/pkg/router"
)

func main() {

	if err := config.Init(); err != nil {
		fmt.Println(err)
	}

	conf := config.GetGlobalConfig()

	// fmt.Println("Web ===>", conf.Web)
	// fmt.Println("DB ===>", conf.Database)
	fmt.Printf("%T", conf)
	fmt.Println()
	fmt.Printf(" %#v", conf)

	r := router.Init()

	r.Run(conf.Web.ServerName + ":" + conf.Web.Port)
}
