package db

import (
	"log"
	"math/rand"
	"slices"

	"github.com/iLopezosa/api-wars/lib/config"
	"github.com/iLopezosa/api-wars/lib/models"
	"github.com/jaswdr/faker"
)

func Reset() {

	err := DBClient.Migrator().DropTable(models.Chat{}, models.Comment{}, models.Message{}, models.Post{}, models.User{}, "participants")
	if err != nil {
		log.Fatal(err)
	}

	err = DBClient.AutoMigrate(models.User{}, models.Post{}, models.Comment{}, models.Message{}, models.Chat{})
	if err != nil {
		log.Fatal(err)
	}
}

func Seed() {
	conf := config.GetConfig()
	fake := faker.New()

	// Start transaction
	tx := DBClient.Begin()

	// Handle panic and rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("SEED ERROR", r)
		}
	}()

	// Create users, posts, and messages
	users := make([]*models.User, conf.NumOfUsers)
	for i := 0; i < conf.NumOfUsers; i++ {

		posts := make([]models.Post, conf.NumOfPosts)
		for j := 0; j < conf.NumOfPosts; j++ { // For each post
			consumerID := uint64(rand.Intn(conf.NumOfUsers)) + 1 // Pick a random consumer
			if consumerID == uint64(i)+1 {                       // If the consumer is the same as the publisher, increment the consumerID
				consumerID = (consumerID % uint64(conf.NumOfUsers)) + 1
			}

			posts[j] = models.Post{ // Create the post
				Title:   fake.Company().CatchPhrase(),
				Content: fake.Lorem().Paragraph(1),
				Comments: []models.Comment{
					{
						Content: "I published this post.",
						UserID:  uint64(i) + 1,
					},
					{
						Content: "I consumed this post.",
						UserID:  consumerID,
					},
				},
			}
		}

		users[i] = &models.User{
			Name:  fake.Person().Name(),
			Email: fake.Internet().Email(),
			Posts: posts,
		}
	}

	// Create users or rollback
	if result := tx.Create(users); result.Error != nil {
		tx.Rollback()
		log.Fatal("SEED ERRROR", result.Error)
	}
	// Commit transaction
	tx.Commit()
	tx = DBClient.Begin()

	// Create chats and messages
	chats := make([]*models.Chat, conf.NumOfChats)
	for i := 0; i < conf.NumOfChats; i++ {
		// Create a slice with all the users IDs to pick from
		unusedUsers := make([]uint64, conf.NumOfUsers)
		for j := 0; j < conf.NumOfUsers; j++ {
			unusedUsers[j] = uint64(j) + 1
		}
		// Pick a random number of participants
		numOfParticipants := rand.Intn(conf.MaxNumOfParticipants) + 1
		participants := make([]*models.User, numOfParticipants)
		for j := 0; j < numOfParticipants; j++ { // For each participant in the chat
			unusedUserIndex := rand.Intn(len(unusedUsers)) // Pick a random user
			participantID := unusedUsers[unusedUserIndex]  // Get the user ID
			participants[j] = &models.User{                // Add the user to the chat
				ID: participantID,
			}
			unusedUsers = slices.Delete(unusedUsers, unusedUserIndex, unusedUserIndex) // Remove the user from the slice
		}
		// Pick a random number of messages
		numOfMessages := rand.Intn(conf.MaxNumOfMessages) + 1
		messages := make([]models.Message, numOfMessages)
		for j := 0; j < numOfMessages; j++ { // For each message in the chat
			usrIndex := uint64(rand.Intn(len(participants))) // Pick a random participant
			messages[j] = models.Message{                    // Add the message to the chat
				Content: fake.Lorem().Sentence(rand.Intn(conf.MaxNumOfWords) + 1),
				UserID:  participants[usrIndex].ID,
			}
		}
		// Create the chat
		chats[i] = &models.Chat{
			Messages:     messages,
			Participants: participants,
		}
	}

	// Create users or rollback
	if result := tx.Create(chats); result.Error != nil {
		tx.Rollback()
		log.Fatal("SEED ERRROR", result.Error)
	}
	// Commit transaction
	tx.Commit()
}
