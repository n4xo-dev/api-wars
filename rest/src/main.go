package main

import (
	"fmt"

	"github.com/iLopezosa/api-wars/rest/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func resetDB(db *gorm.DB) {
	db.Migrator().DropTable(&models.User{}, &models.Post{}, &models.Comment{}, &models.Message{}, &models.Chat{})
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Message{}, &models.Chat{})

}

func main() {
	fmt.Println("Creating connection to the database...")

	// db := dbConnect()
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Madrid"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Resetting the database...")

	resetDB(db)
}
