package tests

import (
	"encoding/json"
	"fmt"

	"github.com/iLopezosa/api-wars/rest/db"
	"github.com/iLopezosa/api-wars/rest/models"
)

func TestMessagesComplete() {
	fmt.Println("\n---TestMessagesComplete---")
	fmt.Println("\n#1 > Creating new message...")

	m := &models.Message{
		ChatID:  1,
		UserID:  1,
		Content: "I am a message",
	}

	if err := db.MessageUpsert(m); err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	m2, err := db.MessageRead(m.ID)
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err := json.MarshalIndent(m2, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	fmt.Printf("%d: %s\n", m.ID, string(b))

	fmt.Println("\n#2 > Updating message...")

	m.Content = "I am an updated message"

	if err = db.MessageUpsert(m); err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	m2, err = db.MessageRead(m.ID)
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(m2, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	fmt.Println(string(b))

	fmt.Println("\n#3 > List all messages:")

	messages, err := db.MessageList()
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(messages, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	fmt.Println(string(b))

	fmt.Println("\n#4 > List all messages by chat id:")

	messages, err = db.MessageListByChatID(1)
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(messages, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	fmt.Println(string(b))

	fmt.Println("\n#5 > List all messages by user id:")

	messages, err = db.MessageListByUserID(1)
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(messages, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	fmt.Println(string(b))

	fmt.Println("\n#6 > List all messages by chat id and user id:")

	messages, err = db.MessageListByChatIDAndUserID(1, 1)
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(messages, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	fmt.Println(string(b))

	fmt.Println("\n#7 > Deleting message...")

	if err = db.MessageDelete(m.ID); err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	if _, err = db.MessageRead(m.ID); err != nil {
		fmt.Printf("\nMessage %d deleted successfully", m.ID)
	} else {
		fmt.Printf("\nTEST ERROR: Message %d not deleted", m.ID)
		return
	}
}
