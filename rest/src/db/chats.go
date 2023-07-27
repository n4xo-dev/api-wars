package db

import (
	"github.com/iLopezosa/api-wars/rest/src/models"
	"gorm.io/gorm"
)

// Updater or creates a chat if the id provided within the chat is found or not, respectively
func ChatUpsert(chat *models.Chat) error {

	ctx := DBClient.Model(&models.Chat{}).Save(chat)

	return ctx.Error
}

// Gets the data of the chat with the provided id
func ChatRead(id uint, eager bool) (models.Chat, error) {

	var chat = models.Chat{
		ID: id,
	}
	var ctx *gorm.DB
	if eager {
		ctx = DBClient.Preload("Participants").Preload("Messages").First(&chat)
	} else {
		ctx = DBClient.First(&chat)
	}

	return chat, ctx.Error
}

// Deletes the chat with the provided id
func ChatDelete(id uint) error {

	var chat = models.Chat{
		ID: id,
	}
	ctx := DBClient.Delete(&chat)

	return ctx.Error
}

// Gets the data of all the chats
func ChatList() ([]models.Chat, error) {

	var chats []models.Chat
	ctx := DBClient.Find(&chats)

	return chats, ctx.Error
}
