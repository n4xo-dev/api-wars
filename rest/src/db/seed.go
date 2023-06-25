package db

import (
	"log"
	"math/rand"
	"time"

	"github.com/iLopezosa/api-wars/rest/src/config"
	"github.com/iLopezosa/api-wars/rest/src/models"
	"github.com/jaswdr/faker"
	"gorm.io/gorm"
)

func Reset(db *gorm.DB) {

	err := db.Migrator().DropTable(models.Chat{}, models.Comment{}, models.Message{}, models.Post{}, models.User{}, "participants")
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(models.User{}, models.Post{}, models.Comment{}, models.Message{}, models.Chat{})
	if err != nil {
		log.Fatal(err)
	}
}

func Seed(db *gorm.DB) {
	conf := config.GetConfig()
	fake := faker.New()

	var users []*models.User
	for i := 1; i <= conf.NumOfUsers; i++ {

		var posts []models.Post
		for j := 1; j <= conf.NumOfPosts; j++ {
			posts = append(posts, models.Post{
				Title:   fake.Company().CatchPhrase(),
				Content: fake.Lorem().Paragraph(1),
				Comments: []models.Comment{
					{
						Content: "I published this post.",
						UserID:  uint(i),
					},
					{
						Content: "I consumed this post.",
						UserID:  uint(rand.Intn(conf.NumOfUsers)) + 1,
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

	db.Create(users)

	var chats []*models.Chat
	for i := 1; i <= conf.NumOfChats; i++ {

		var participants []*models.User
		numOfParticipants := rand.Intn(conf.MaxNumOfParticipants) + 1
		for j := 1; j <= numOfParticipants; j++ {
			participants = append(participants, &models.User{
				ID: uint(rand.Intn(conf.NumOfUsers) + 1),
			})
		}

		var messages []models.Message
		numOfMessages := rand.Intn(conf.MaxNumOfMessages) + 1
		for j := 1; j <= numOfMessages; j++ {
			messages = append(messages, models.Message{
				Content:   fake.Lorem().Sentence(j),
				Timestamp: fake.Time().Time(time.Now()),
				UserID:    uint(rand.Intn(conf.NumOfUsers)) + 1,
			})
		}

		chats = append(chats, &models.Chat{
			ID:           uint(i),
			Messages:     messages,
			Participants: participants,
		})
	}

	db.Create(chats)
}
