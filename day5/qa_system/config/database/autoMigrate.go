package database

import (
	"day5/qa_system/app/models"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Question{},
		&models.NameMap{},
		&models.Submit{})
}
