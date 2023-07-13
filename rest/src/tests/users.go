package tests

import (
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/iLopezosa/api-wars/rest/src/db"
	"github.com/iLopezosa/api-wars/rest/src/models"
)

func TestUsersComplete() {
	fmt.Println("Creating new user...")

	u := &models.User{
		Name:  "Mr. Han",
		Email: "mr.han@lnkn.com",
	}

	if err := db.UserUpsert(u); err != nil {
		log.Fatal(err)
	}

	u2, err := db.UserRead(u.ID)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(u2)
	time.Sleep(5 * time.Second)
	fmt.Println("Updating user...")

	u.Email = "mr.han@linkin.park"

	if err = db.UserUpsert(u); err != nil {
		log.Fatal(err)
	}

	u2, err = db.UserRead(u.ID)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(u2)
	time.Sleep(5 * time.Second)
	fmt.Println("List all users:")

	users, err := db.UserList()
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(users)
	time.Sleep(5 * time.Second)
	fmt.Println("Find user with email: mr.han@linkin.park")

	users, err = db.UserFindByEmail("mr.han@linkin.park")
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(users)
	time.Sleep(5 * time.Second)
	fmt.Println("Deleting new user...")

	if err = db.UserDelete(u.ID); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE!")
}
