package tests

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/iLopezosa/api-wars/rest/src/db"
	"github.com/iLopezosa/api-wars/rest/src/models"
)

func TestChatsComplete() {
	fmt.Println("\n---TestChatsComplete---")
	fmt.Println("\n#1 > Creating new chat...")

	u1, err := db.UserRead(1)
	if err != nil {
		log.Fatal(err)
	}
	u2, err := db.UserRead(2)
	if err != nil {
		log.Fatal(err)
	}

	c := &models.Chat{
		Participants: []*models.User{&u1, &u2},
	}
	fmt.Printf("%+v\n", c)
	if err := db.ChatUpsert(c); err != nil {
		log.Fatal(err)
	}

	c2, err := db.ChatRead(c.ID, true)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(c2, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d: %s\n", c.ID, string(b))
}
