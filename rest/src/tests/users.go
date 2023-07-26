package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/iLopezosa/api-wars/rest/src/db"
	"github.com/iLopezosa/api-wars/rest/src/models"
)

func TestUsersComplete() {
	fmt.Println("\n---TestUsersComplete---")
	fmt.Println("\n#1 > Creating new user...")

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

	b, err := json.MarshalIndent(u2, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d: %s\n", u.ID, string(b))
	time.Sleep(5 * time.Second)
	fmt.Println("\n#2 > Updating user...")

	u.Email = "mr.han@linkin.park"

	if err = db.UserUpsert(u); err != nil {
		log.Fatal(err)
	}

	u2, err = db.UserRead(u.ID)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(u2, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	time.Sleep(5 * time.Second)
	fmt.Println("\n#3 > List all users:")

	users, err := db.UserList()
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	time.Sleep(5 * time.Second)
	fmt.Println("\n#4 > Find user with email: mr.han@linkin.park")

	users, err = db.UserFindByEmail("mr.han@linkin.park")
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	time.Sleep(5 * time.Second)
	fmt.Println("\n#5 > Deleting new user...")

	if err = db.UserDelete(u.ID); err != nil {
		log.Fatal(err)
	}

	_, err = db.UserRead(u.ID)
	if err != nil {
		fmt.Println("User deleted")
	} else {
		log.Fatal("User not deleted")
	}

	fmt.Println("DONE!")
}
