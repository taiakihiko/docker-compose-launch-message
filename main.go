package main

import (
	"time"
    "github.com/gin-gonic/gin"
)

func main() {
	
	time.Sleep(time.Second * 10) // 起動に時間がかかるサーバー
	
	router := gin.Default()
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H {"message": "server is ready"})
	})
	router.Run()
}
