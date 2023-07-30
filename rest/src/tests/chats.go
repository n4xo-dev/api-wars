package tests

import (
	"encoding/json"
	"fmt"

	"github.com/iLopezosa/api-wars/rest/src/db"
	"github.com/iLopezosa/api-wars/rest/src/models"
)

func TestChatsComplete() {
	fmt.Println("\n---TestChatsComplete---")
	fmt.Println("\n#1 > Creating new chat...")

	u1 := models.User{
		ID: 1,
	}
	u2 := models.User{
		ID: 2,
	}

	c := &models.Chat{
		Participants: []*models.User{&u1, &u2},
	}
	fmt.Printf("%+v\n", c)
	if err := db.ChatUpsert(c); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	c2, err := db.ChatRead(c.ID, true)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	b, err := json.MarshalIndent(c2, "", "  ")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Printf("%d: %s\n", c.ID, string(b))

	fmt.Println("\n#2 > Adding message to chat...")
	m := models.Message{
		ChatID:  c.ID,
		UserID:  u1.ID,
		Content: "Hello there!",
	}
	c.Messages = append(c.Messages, m)
	if err := db.ChatUpsert(c); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	c2, err = db.ChatRead(c.ID, true)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(c2, "", "  ")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(b))

	fmt.Println("\n#3 > List all chats:")
	chats, err := db.ChatList(true)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(chats, "", "  ")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(b))

	fmt.Println("\n#4 > Deleting chat...")
	if err := db.ChatDelete(c.ID); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	_, err = db.ChatRead(c.ID, false)
	if err != nil {
		fmt.Printf("\nChat %d deleted successfully\n", c.ID)
	} else {
		fmt.Printf("\nERROR: Chat %d still exists", c.ID)
		return
	}

	fmt.Println("DONE!")
}
