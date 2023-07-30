package db

import (
	"github.com/iLopezosa/api-wars/rest/src/models"
	"gorm.io/gorm"
)

// Updater or creates a chat if the id provided within the chat is found or not, respectively
func ChatUpsert(chat *models.Chat) error {

	ctx := DBClient.Save(chat)

	return ctx.Error
}

// Gets the data of the chat with the provided id
func ChatRead(id uint64, eager bool) (models.ChatDTO, error) {

	var chat = models.ChatDTO{}
	var ctx *gorm.DB
	if eager {
		ctx = DBClient.Model(&models.Chat{}).Preload("Participants").Preload("Messages").First(&chat, id)
	} else {
		ctx = DBClient.Model(&models.Chat{}).First(&chat, id)
	}

	return chat, ctx.Error
}

// Deletes the chat with the provided id
func ChatDelete(id uint64) error {

	var chat = models.Chat{
		ID: id,
	}
	ctx := DBClient.Delete(&chat)

	return ctx.Error
}

// Gets the data of all the chats
func ChatList(eager bool) ([]models.ChatDTO, error) {

	var chats []models.ChatDTO

	var ctx *gorm.DB
	if eager {
		ctx = DBClient.Model(&models.Chat{}).Preload("Participants").Preload("Messages").Find(&chats)
	} else {
		ctx = DBClient.Model(&models.Chat{}).Find(&chats)
	}

	return chats, ctx.Error
}
