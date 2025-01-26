package migrations

import (
	"server/db"
	"server/models"
)

func Migrate() {
	db.DB.AutoMigrate(&models.Users{})
	db.DB.AutoMigrate(&models.UserProfiles{})
	db.DB.AutoMigrate(&models.WeightLog{})
}
