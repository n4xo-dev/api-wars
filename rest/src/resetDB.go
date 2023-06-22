package main

import (
	"github.com/iLopezosa/api-wars/rest/src/models"
	"gorm.io/gorm"
)

func resetDB(db *gorm.DB) {
	db.Migrator().DropTable(&models.User{}, &models.Post{}, &models.Comment{}, &models.Message{}, &models.Chat{})
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Message{}, &models.Chat{})

}
