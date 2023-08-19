package tests

import (
	"encoding/json"
	"fmt"

	"github.com/iLopezosa/api-wars/rest/db"
	"github.com/iLopezosa/api-wars/rest/models"
)

func TestUsersComplete() {
	fmt.Println("\n---TestUsersComplete---")
	fmt.Println("\n#1 > Creating new user...")

	u := &models.User{
		Name:  "Mr. Han",
		Email: "mr.han@lnkn.com",
	}

	if err := db.UserUpsert(u); err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	u2, err := db.UserRead(u.ID)
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err := json.MarshalIndent(u2, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	fmt.Printf("%d: %s\n", u.ID, string(b))
	fmt.Println("\n#2 > Updating user...")

	u.Email = "mr.han@linkin.park"

	if err = db.UserUpsert(u); err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	u2, err = db.UserRead(u.ID)
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(u2, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println("\n#3 > List all users:")

	users, err := db.UserList()
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println("\n#4 > Find user with email: mr.han@linkin.park")

	users, err = db.UserFindByEmail("mr.han@linkin.park")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println("\n#5 > Deleting new user...")

	if err = db.UserDelete(u.ID); err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	_, err = db.UserRead(u.ID)
	if err != nil {
		fmt.Printf("\nUser %d deleted successfully\n", u.ID)
	} else {
		fmt.Printf("\nTEST ERROR: User %d not deleted\n", u.ID)
		return
	}

	fmt.Println("DONE!")
}
