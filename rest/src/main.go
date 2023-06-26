package main

import (
	"fmt"
	"log"
	"time"

	"github.com/iLopezosa/api-wars/rest/src/db"
	"github.com/iLopezosa/api-wars/rest/src/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Creating connection to the database...")

	db.Connect()

	fmt.Println("Resetting the database...")

	db.Reset()

	fmt.Println("Seeding database...")

	db.Seed()

	fmt.Println("Creating new user...")

	u := &models.User{
		Name:  "Mr. Han",
		Email: "mr.han@lnkn.com",
	}

	if err = db.UserUpsert(u); err != nil {
		log.Fatal(err)
	}

	u2, err := db.UserRead(u.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", u2)
	time.Sleep(10 * time.Second)
	fmt.Println("Updating user...")

	u.Email = "mr.han@linkin.park"

	if err = db.UserUpsert(u); err != nil {
		log.Fatal(err)
	}

	u2, err = db.UserRead(u.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", u2)
	time.Sleep(10 * time.Second)
	fmt.Println("Deleting user...")

	if err = db.UserDelete(u.ID); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE!")
}
