package db

import (
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBClient *gorm.DB

var gormOnce sync.Once

func Connect() *gorm.DB {
	gormOnce.Do(func() {
		dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Madrid"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			PrepareStmt: true,
			Logger:      logger.Default.LogMode(logger.Silent),
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
