package migrations

import (
	"server/db"
	"server/models"
)

func Migrate() {
	db.DB.AutoMigrate(&models.Users{})
}
