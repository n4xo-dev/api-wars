package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/iLopezosa/api-wars/rest/src/db"
	"github.com/iLopezosa/api-wars/rest/src/models"
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
		log.Fatal(err)
	}

	m2, err := db.MessageRead(m.ID)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(m2, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d: %s\n", m.ID, string(b))
	time.Sleep(5 * time.Second)

	fmt.Println("\n#2 > Updating message...")

	m.Content = "I am an updated message"

	if err = db.MessageUpsert(m); err != nil {
		log.Fatal(err)
	}

	m2, err = db.MessageRead(m.ID)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(m2, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
	time.Sleep(5 * time.Second)

	fmt.Println("\n#3 > List all messages:")

	messages, err := db.MessageList()
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(messages, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
	time.Sleep(5 * time.Second)

	fmt.Println("\n#4 > List all messages by chat id:")

	messages, err = db.MessageListByChatID(1)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(messages, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
	time.Sleep(5 * time.Second)

	fmt.Println("\n#5 > List all messages by user id:")

	messages, err = db.MessageListByUserID(1)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(messages, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
	time.Sleep(5 * time.Second)

	fmt.Println("\n#6 > List all messages by chat id and user id:")

	messages, err = db.MessageListByChatIDAndUserID(1, 1)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(messages, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
	time.Sleep(5 * time.Second)

	fmt.Println("\n#7 > Deleting message...")

	if err = db.MessageDelete(m.ID); err != nil {
		log.Fatal(err)
	}

	if _, err = db.MessageRead(m.ID); err != nil {
		fmt.Println("Message deleted successfully")
	} else {
		log.Fatal("Message not deleted")
	}
}
