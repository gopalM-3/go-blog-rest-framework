package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gopalM-3/go-blog-rest-framework/models"
)

func PostBlog(context *gin.Context) {
	// reading request body
	var body struct {
		Category string
		Title 	 string
		Excerpt  string
		Content  string
		Slug 	 string
		Author 	 string
		Hashtags string
		Status 	 bool
		Upvotes  int
	}

	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body"})
		return
	}

	// creating the blog
	blog := models.Blog{
		Category: body.Category, 
		Title: body.Title, 
		Excerpt: body.Excerpt, 
		Content: body.Content, 
		Slug: body.Slug, 
		Author: , 
		Hashtags: body.Hashtags, 
		Status: body.Status, 
		Upvotes: 0,
	}
}