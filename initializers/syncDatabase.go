package initializers

import "github.com/PongthepNuchwet/go-jwt/model"

func SyncDatabase() {
	DB.AutoMigrate(&model.User{})
}
