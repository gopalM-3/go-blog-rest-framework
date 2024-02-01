package initializers

import (
	"fmt"

	"github.com/gopalM-3/go-blog-rest-framework/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database migrations made!")
}