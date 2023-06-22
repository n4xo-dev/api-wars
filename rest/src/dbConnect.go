package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect to the database
func dbConnect() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Madrid"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
