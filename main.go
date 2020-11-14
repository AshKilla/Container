package main

import (
	"github.com/AshKilla/cloud-join-you/router"
	"github.com/AshKilla/cloud-join-you/util"
)

func main() {
	ginApp := router.GetApp()
	if err := ginApp.Run(util.GetPort()); err != nil {
		panic(err)
	}
}
