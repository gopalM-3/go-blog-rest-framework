package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gopalM-3/go-blog-rest-framework/initializers"
	"github.com/gopalM-3/go-blog-rest-framework/models"
)

func RequireAuth(context *gin.Context) {
	// retrieving the cookie from the request
	tokenString, err := context.Cookie("Authorization")

	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Access denied, login to access"})
		return
	}

	// verifying the extracted token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
	
		return []byte(os.Getenv("HMAC_SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			context.AbortWithStatus(http.StatusUnauthorized)
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Token expired, login again"})
			return
		}

		// fetching user
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			context.AbortWithStatus(http.StatusUnauthorized)
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Access denied, user not found"})
			return
		}

		// attaching the user to the request
		context.Set("user", user)

		context.Next()
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid claims"})
		return
	}
}