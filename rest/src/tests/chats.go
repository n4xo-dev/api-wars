package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

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
		Participants: []*models.User{models.ParseUserDTO(&u1), models.ParseUserDTO(&u2)},
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
	time.Sleep(5 * time.Second)

	fmt.Println("\n#2 > Adding message to chat...")
	m := models.Message{
		ChatID:  c.ID,
		UserID:  u1.ID,
		Content: "Hello there!",
	}
	c.Messages = append(c.Messages, m)
	if err := db.ChatUpsert(c); err != nil {
		log.Fatal(err)
	}

	c2, err = db.ChatRead(c.ID, true)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(c2, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	time.Sleep(5 * time.Second)

	fmt.Println("\n#3 > List all chats:")
	chats, err := db.ChatList(true)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(chats, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	time.Sleep(5 * time.Second)

	fmt.Println("\n#4 > Deleting chat...")
	if err := db.ChatDelete(c.ID); err != nil {
		log.Fatal(err)
	}

	_, err = db.ChatRead(c.ID, false)
	if err != nil {
		fmt.Println("Chat deleted successfully")
	} else {
		log.Fatal("Chat still exists")
	}

	fmt.Println("DONE!")
}
