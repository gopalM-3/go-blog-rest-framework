package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gopalM-3/go-blog-rest-framework/initializers"
	"github.com/gopalM-3/go-blog-rest-framework/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(context *gin.Context) {
	// reading request body
	var body struct {
		Username string
		Email 	 string
		Password string
	}

	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body"})
		return
	}

	// generating hashed password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to hash password"})
		return
	}

	// creating user
	user := models.User{Username: body.Username, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User created successfully!"})
}

func Login(context *gin.Context) {
	// reading request body
	var body struct {
		Email 	 string
		Password string
	}

	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body"})
		return
	}

	// fetching user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email ID"})
		return
	}

	// comparing passwords
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid password"})
		return
	}

	// generating token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// encoding token
	tokenString, err := token.SignedString([]byte(os.Getenv("HMAC_SECRET")))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Token creation failed"})
		return
	}

	// setting the token in the cookie
	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("Authorization", tokenString, int(time.Hour) * 24, "", "", false, true)

	context.JSON(http.StatusOK, gin.H{"message": "User logged in successfully!"})
}

func Verify(context *gin.Context) {
	user, _ := context.Get("user")

	context.JSON(http.StatusOK, gin.H{"message": "Valid token!", "user": user})
}

func Logout(context *gin.Context) {
	// unsetting the Authorization cookie
	context.SetCookie("Authorization", "", -1, "", "", false, true)

	context.JSON(http.StatusOK, gin.H{"message": "User logged out successfully!"})
}