package main

import (
	"fmt"
	"log"

	"github.com/iLopezosa/api-wars/rest/src/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Creating connection to the database...")

	dbClient := db.Connect()

	fmt.Println("Resetting the database...")

	db.Reset(dbClient)

	fmt.Println("Seeding database...")

	db.Seed(dbClient)
}
