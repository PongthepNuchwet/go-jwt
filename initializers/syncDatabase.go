package initializers

import "github.com/PongthepNuchwet/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
