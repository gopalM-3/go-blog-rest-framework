package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gopalM-3/go-blog-rest-framework/controllers"
	"github.com/gopalM-3/go-blog-rest-framework/initializers"
	"github.com/gopalM-3/go-blog-rest-framework/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/verify-token", middleware.RequireAuth, controllers.Verify)
	router.GET("/logout", controllers.Logout)

	router.Run("localhost:8000")
}