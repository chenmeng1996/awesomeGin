package main

import (
	"awesomeGin/biz/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", handler.Ping)
	_ = router.Run(":20000")
}
