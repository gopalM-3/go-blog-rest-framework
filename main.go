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

	router.GET("/", controllers.Homepage)

	router.POST("/user/signup", controllers.Signup)
	router.POST("/user/login", controllers.Login)
	router.GET("/user/logout", controllers.Logout)
	router.GET("/verify-token", middleware.RequireAuth, controllers.Verify)

	router.GET("/blogs/all", controllers.AllBlogs)
	router.GET("/blogs/id/:id", controllers.BlogById)
	router.GET("/blogs/slug/:slug", controllers.BlogBySlug)
	router.GET("/blogs/category/:category", controllers.BlogByCategory)
	router.GET("/blogs/hashtag/:hashtag", controllers.BlogByHashtag)
	router.GET("/blogs/user", middleware.RequireAuth, controllers.BlogByUser)
	router.POST("/blogs/post", middleware.RequireAuth, controllers.PostBlog)
	router.PUT("/blogs/update/:id", middleware.RequireAuth, controllers.UpdateBlog)
	router.DELETE("/blogs/delete/:id", middleware.RequireAuth, controllers.DeleteBlog)

	router.Run("localhost:8000")
}