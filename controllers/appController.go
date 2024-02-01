package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Homepage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Welcome to the Blogger app!"})
}