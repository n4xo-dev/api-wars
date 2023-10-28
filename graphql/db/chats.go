package db

import (
	"github.com/iLopezosa/api-wars/lib/models"
	"gorm.io/gorm"
)

// Updater or creates a chat if the id provided within the chat is found or not, respectively
func ChatUpsert(chat *models.Chat) error {

	ctx := DBClient.Save(chat).Take(chat)

	return ctx.Error
}

// Gets the data of the chat with the provided id
func ChatRead(id uint64) (models.Chat, error) {

	var chat = models.Chat{}
	var ctx *gorm.DB
	ctx = DBClient.Model(&models.Chat{}).First(&chat, id)

	return chat, ctx.Error
}

// Get the messages of the chat with the provided id
func ChatMessages(id uint64) ([]*models.Message, error) {

	var messages []*models.Message
	ctx := DBClient.Model(&models.Message{}).Where("chat_id = ?", id).Find(&messages)

	return messages, ctx.Error
}

// Get the participants of the chat with the provided id
func ChatParticipants(id uint64) ([]*models.User, error) {

	chat := models.Chat{
		ID: id,
	}
	ctx := DBClient.Preload("Participants").Find(&chat)

	return chat.Participants, ctx.Error
}

// Patch update the chat with the provided id
func ChatPatch(chat *models.Chat) error {

	ctx := DBClient.Updates(chat).Take(chat)

	return ctx.Error
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
func ChatList() ([]*models.Chat, error) {

	var chats []*models.Chat

	var ctx *gorm.DB

	ctx = DBClient.Model(&models.Chat{}).Find(&chats)

	return chats, ctx.Error
}
