package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gopalM-3/go-blog-rest-framework/initializers"
	"github.com/gopalM-3/go-blog-rest-framework/middleware"
	"github.com/gopalM-3/go-blog-rest-framework/models"
)

func AllBlogs(context *gin.Context) {
	var blogs []models.Blog
	result := initializers.DB.Find(&blogs)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to retrieve blogs"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

func BlogById(context *gin.Context) {
	id := context.Param("id")

	var blog models.Blog
	result := initializers.DB.First(&blog, id)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to retrieve blog"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"blog": blog})
}

func BlogBySlug(context *gin.Context) {
	slug := context.Param("slug")

	var blog models.Blog
	result := initializers.DB.First(&blog, "slug = ?", slug)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to retrieve blog"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"blog": blog})
}

func BlogByCategory(context *gin.Context) {
	category := context.Param("category")

	var blogs []models.Blog
	result := initializers.DB.Find(&blogs, "category = ?", category)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to retrieve blogs"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

func BlogByHashtag(context *gin.Context) {
	hashtag := context.Param("hashtag")

	var blogs []models.Blog
	searchKey := "%" + hashtag + "%"
	result := initializers.DB.Where("hashtags LIKE ?", searchKey).Find(&blogs)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to retrieve blogs"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

func BlogByUser(context *gin.Context) {
	claims := middleware.ExtractClaims(context)

	var blogs []models.Blog
	result := initializers.DB.Where("author = ?", claims["username"]).Find(&blogs)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to retrieve blogs"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

func PostBlog(context *gin.Context) {
	// reading request body
	var body struct {
		Category string
		Title 	 string
		Excerpt  string
		Content  string
		Slug 	 string
		Hashtags string
		Status 	 bool
	}

	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body"})
		return
	}

	// creating the blog
	claims := middleware.ExtractClaims(context)

	blog := models.Blog{
		Category: body.Category,
		Title:    body.Title,
		Excerpt:  body.Excerpt,
		Content:  body.Content,
		Slug:     body.Slug,
		Author:   claims["username"].(string),
		Hashtags: body.Hashtags,
		Status:   body.Status,
		Votes:    0,
	}
	result := initializers.DB.Create(&blog)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to post blog"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Blog posted successfully!", "blog": blog})
}

func UpdateBlog(context *gin.Context) {
	id := context.Param("id")
	claims := middleware.ExtractClaims(context)

	var blog models.Blog
	result := initializers.DB.First(&blog, id)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to retrieve blog"})
		return
	}

	// verifying if the logged in user is the author of the queried id
	if blog.Author == claims["username"] {
		// reading request body
		var body struct {
			Category string
			Title 	 string
			Excerpt  string
			Content  string
			Slug 	 string
			Hashtags string
			Status 	 bool
		}
	
		if context.Bind(&body) != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body"})
			return
		}

		// updating blog
		updatedBlog := models.Blog{
			Category: blog.Category,
			Title:    body.Title,
			Excerpt:  body.Excerpt,
			Content:  body.Content,
			Slug:     body.Slug,
			Author:   blog.Author,
			Hashtags: body.Hashtags,
			Status:   body.Status,
			Votes:    blog.Votes,
		}

		initializers.DB.Model(&blog).Updates(updatedBlog)
		initializers.DB.Save(&blog)

		result := initializers.DB.First(&blog, id)
	
		if result.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to retrieve blog"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully!", "blog": blog})
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized, user footprint mismatch"})
		return
	}
}

func DeleteBlog(context *gin.Context) {
	id := context.Param("id")
	claims := middleware.ExtractClaims(context)

	var blog models.Blog
	result := initializers.DB.First(&blog, id)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to retrieve blog"})
		return
	}

	// verifying if the logged in user is the author of the queried id
	if blog.Author == claims["username"] {
		initializers.DB.Delete(&blog)

		context.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully!"})
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized, user footprint mismatch"})
		return
	}
}