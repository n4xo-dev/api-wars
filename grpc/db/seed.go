package db

import (
	"log"
	"math/rand"

	"github.com/iLopezosa/api-wars/grpc/config"
	"github.com/iLopezosa/api-wars/grpc/models"
	"github.com/jaswdr/faker"
	"golang.org/x/exp/slices"
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

	// Create users, posts and comments
	var users []*models.User
	for i := 1; i <= conf.NumOfUsers; i++ {

		var posts []models.Post
		for j := 1; j <= conf.NumOfPosts; j++ {
			consumerID := uint64(rand.Intn(conf.NumOfUsers)) + 1
			// If the consumer is the same as the publisher, increment the consumerID
			if consumerID == uint64(i) {
				consumerID = (consumerID % uint64(conf.NumOfUsers)) + 1
			}

			posts = append(posts, models.Post{
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
			})
		}

		users = append(users, &models.User{
			Name:  fake.Person().Name(),
			Email: fake.Internet().Email(),
			Posts: posts,
		})
	}

	// Create users or rollback
	if result := tx.Create(users); result.Error != nil {
		tx.Rollback()
		log.Fatal(result.Error)
	}

	// Create chats and messages
	var chats []*models.Chat
	for i := 1; i <= conf.NumOfChats; i++ {

		var participants []*models.User
		numOfParticipants := rand.Intn(conf.MaxNumOfParticipants) + 1
		for j := 1; j <= numOfParticipants; j++ {
			participantID := uint64(rand.Intn(conf.NumOfUsers)) + 1
			// If the participant is already in the chat, increment the participantID
			if slices.ContainsFunc(participants, func(p *models.User) bool { return p.ID == participantID }) {
				participantID = (participantID % uint64(conf.NumOfUsers)) + 1
			}
			participants = append(participants, &models.User{
				ID: participantID,
			})
		}

		var messages []models.Message
		numOfMessages := rand.Intn(conf.MaxNumOfMessages) + 1
		for j := 1; j <= numOfMessages; j++ {
			usrIndex := uint64(rand.Intn(len(participants)))
			messages = append(messages, models.Message{
				Content: fake.Lorem().Sentence(j),
				UserID:  participants[usrIndex].ID,
			})
		}

		chats = append(chats, &models.Chat{
			Messages:     messages,
			Participants: participants,
		})
	}

	// Create chats or rollback
	if result := tx.Create(chats); result.Error != nil {
		tx.Rollback()
		log.Fatal(result.Error)
	}

	// Commit transaction
	tx.Commit()
}
