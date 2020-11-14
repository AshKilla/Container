package router

import (
	"github.com/AshKilla/cloud-join-you/controller"
	"github.com/gin-gonic/gin"
)

func GetApp() (router *gin.Engine) {
	router = gin.Default()
	router.GET("/ping", controller.Ping)
	router.GET("/", controller.Index)
	return
}
