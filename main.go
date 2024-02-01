package main

import (

	"github.com/gin-gonic/gin"
	"github.com/gopalM-3/go-blog-rest-framework/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run("localhost:8000")
}