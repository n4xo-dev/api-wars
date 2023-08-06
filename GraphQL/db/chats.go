package db

import (
	"github.com/iLopezosa/api-wars/graphql/graph/model"
	"gorm.io/gorm"
)

// Updater or creates a chat if the id provided within the chat is found or not, respectively
func ChatUpsert(chat *model.Chat) error {

	ctx := DBClient.Save(chat).Take(chat)

	return ctx.Error
}

// Gets the data of the chat with the provided id
func ChatRead(id uint64, eager bool) (model.Chat, error) {

	var chat = model.Chat{}
	var ctx *gorm.DB
	if eager {
		ctx = DBClient.Model(&model.Chat{}).Preload("Participants").Preload("Messages").First(&chat, id)
	} else {
		ctx = DBClient.Model(&model.Chat{}).First(&chat, id)
	}

	return chat, ctx.Error
}

// Patch update the chat with the provided id
func ChatPatch(chat *model.Chat) error {

	ctx := DBClient.Updates(chat).Take(chat)

	return ctx.Error
}

// Deletes the chat with the provided id
func ChatDelete(id uint64) error {

	var chat = model.Chat{
		ID: id,
	}
	ctx := DBClient.Delete(&chat)

	return ctx.Error
}

// Gets the data of all the chats
func ChatList(eager bool) ([]model.Chat, error) {

	var chats []model.Chat

	var ctx *gorm.DB
	if eager {
		ctx = DBClient.Model(&model.Chat{}).Preload("Participants").Preload("Messages").Find(&chats)
	} else {
		ctx = DBClient.Model(&model.Chat{}).Find(&chats)
	}

	return chats, ctx.Error
}
