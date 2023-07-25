package db

import (
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

var gormOnce sync.Once

func Connect() *gorm.DB {
	gormOnce.Do(func() {
		dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Madrid"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			PrepareStmt: true,
		})

		if err != nil {
			log.Fatal(err)
		}

		DBClient = db
	})

	return DBClient
}

func Disconnect() {
	sqlDB, err := DBClient.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.Close()
}
