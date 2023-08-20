package db

import (
	"log"
	"math/rand"
	"slices"

	"github.com/iLopezosa/api-wars/rest/config"
	"github.com/iLopezosa/api-wars/rest/models"
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
		}
	}()

	// Create users, posts, and messages
	users := make([]*models.User, conf.NumOfUsers)
	for i := 1; i <= conf.NumOfUsers; i++ {

		posts := make([]models.Post, conf.NumOfPosts)
		for j := 1; j <= conf.NumOfPosts; j++ {
			consumerID := uint64(rand.Intn(conf.NumOfUsers)) + 1
			// If the consumer is the same as the publisher, increment the consumerID
			if consumerID == uint64(i) {
				consumerID = (consumerID % uint64(conf.NumOfUsers)) + 1
			}

			posts[j-1] = models.Post{
				Title:   fake.Company().CatchPhrase(),
				Content: fake.Lorem().Paragraph(1),
				Comments: []models.Comment{
					{
						Content: "I published this post.",
						UserID:  uint64(i),
					},
					{
						Content: "I consumed this post.",
						UserID:  consumerID,
					},
				},
			}
		}

		users[i-1] = &models.User{
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

	// Create chats and messages
	chats := make([]*models.Chat, conf.NumOfChats)
	for i := 1; i <= conf.NumOfChats; i++ {

		numOfParticipants := rand.Intn(conf.MaxNumOfParticipants) + 1
		participants := make([]*models.User, numOfParticipants)
		for j := 1; j <= numOfParticipants; j++ {
			participantID := uint64(rand.Intn(conf.NumOfUsers)) + 1
			// If the participant is already in the chat, increment the participantID
			if slices.ContainsFunc(participants, func(p *models.User) bool { return p.ID == participantID }) {
				participantID = (participantID % uint64(conf.NumOfUsers))
			}
			participants[j-1] = &models.User{
				ID: participantID,
			}
		}

		messages := make([]models.Message, conf.MaxNumOfMessages)
		numOfMessages := rand.Intn(conf.MaxNumOfMessages) + 1
		for j := 1; j <= numOfMessages; j++ {
			usrIndex := uint64(rand.Intn(len(participants)))
			messages[j-1] = models.Message{
				Content: fake.Lorem().Sentence(j),
				UserID:  participants[usrIndex].ID,
			}
		}

		chats[i-1] = &models.Chat{
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
