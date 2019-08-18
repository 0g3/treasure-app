package main

import (
	"github.com/0g3/treasure-app/app/config"
	"github.com/gin-gonic/gin"
	"log"
)

var router = gin.Default()

func init() {
	router.GET("/alive", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I'm OK!",
		})
	})
}

func main() {
	log.Fatal(router.Run(":" + config.Env().ServerPort))
}
