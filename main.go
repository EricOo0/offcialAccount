package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"wxcloudrun-golang/service"
)

func main() {
	http := gin.Default()
	publicRouter := http.Group("")
	{

		publicRouter.GET("/send_wx_msg", service.VerifyWxToken())
		publicRouter.POST("/send_wx_msg", service.ProcessWxMsgHandler())

	}
	if err := http.Run(":80"); err != nil {
		log.Fatalln("http server is fail!")

	}
}
