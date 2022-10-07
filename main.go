package main

import (
	"fmt"

	"github.com/dzgntutar/ccgin/pkg/config"
	"github.com/dzgntutar/ccgin/pkg/router"
)

func main() {

	if err := config.Init(); err != nil {
		//log
		fmt.Println(err)
	}

	conf := config.GetGlobalConfig()

	fmt.Println("Config is ready  --->")
	fmt.Printf("%+v\n", conf)

	r := router.Init()

	r.Run()
}
