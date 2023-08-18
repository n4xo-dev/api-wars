package db

import (
	"log"
	"math/rand"
	"time"

	"github.com/iLopezosa/api-wars/graphql/config"
	"github.com/iLopezosa/api-wars/graphql/graph/model"
	"github.com/jaswdr/faker"
	"golang.org/x/exp/slices"
)

func Reset() {

	err := DBClient.Migrator().DropTable(model.Chat{}, model.Comment{}, model.Message{}, model.Post{}, model.User{}, "participants")
	if err != nil {
		log.Fatal(err)
	}

	err = DBClient.AutoMigrate(model.User{}, model.Post{}, model.Comment{}, model.Message{}, model.Chat{})
	if err != nil {
		log.Fatal(err)
	}
}

func Seed() {
	conf := config.GetConfig()
	fake := faker.New()

	var users []*model.User
	for i := 1; i <= conf.NumOfUsers; i++ {

		var posts []model.Post
		for j := 1; j <= conf.NumOfPosts; j++ {
			consumerID := uint64(rand.Intn(conf.NumOfUsers)) + 1
			// If the consumer is the same as the publisher, increment the consumerID
			if consumerID == uint64(i) {
				consumerID++
			}

			posts = append(posts, model.Post{
				Title:   fake.Company().CatchPhrase(),
				Content: fake.Lorem().Paragraph(1),
				Comments: []model.Comment{
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

		users = append(users, &model.User{
			Name:  fake.Person().Name(),
			Email: fake.Internet().Email(),
			Posts: posts,
		})
	}

	DBClient.Create(users)

	var chats []*model.Chat
	for i := 1; i <= conf.NumOfChats; i++ {

		var participants []*model.User
		numOfParticipants := rand.Intn(conf.MaxNumOfParticipants) + 1
		for j := 1; j <= numOfParticipants; j++ {
			participantID := uint64(rand.Intn(conf.NumOfUsers) + 1)

			// If the participant is already in the chat, increment the participantID
			if slices.ContainsFunc(participants, func(p *model.User) bool { return p.ID == participantID }) {
				participantID++
			}
			participants = append(participants, &model.User{
				ID: participantID,
			})
		}

		var messages []model.Message
		numOfMessages := rand.Intn(conf.MaxNumOfMessages) + 1
		for j := 1; j <= numOfMessages; j++ {
			usrIndex := uint64(rand.Intn(len(participants)))
			messages = append(messages, model.Message{
				Content: fake.Lorem().Sentence(j),
				UserID:  participants[usrIndex].ID,
			})
		}

		chats = append(chats, &model.Chat{
			Messages:     messages,
			Participants: participants,
		})
	}

	DBClient.Create(chats)

	// Ensure that the chats are saved in the database
	time.Sleep(5 * time.Second)
}
