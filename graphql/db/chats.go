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
func ChatRead(id uint64) (model.Chat, error) {

	var chat = model.Chat{}
	var ctx *gorm.DB
	ctx = DBClient.Model(&model.Chat{}).First(&chat, id)

	return chat, ctx.Error
}

// Get the messages of the chat with the provided id
func ChatMessages(id uint64) ([]*model.Message, error) {

	var messages []*model.Message
	ctx := DBClient.Model(&model.Message{}).Where("chat_id = ?", id).Find(&messages)

	return messages, ctx.Error
}

// Get the participants of the chat with the provided id
func ChatParticipants(id uint64) ([]*model.User, error) {

	chat := model.Chat{
		ID: id,
	}
	ctx := DBClient.Preload("Participants").Find(&chat)

	return chat.Participants, ctx.Error
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
func ChatList() ([]*model.Chat, error) {

	var chats []*model.Chat

	var ctx *gorm.DB

	ctx = DBClient.Model(&model.Chat{}).Find(&chats)

	return chats, ctx.Error
}

// Gets the data of the chats with the provided user id
func ChatListByUserID(userID uint64) ([]*model.Chat, error) {

	user := model.User{
		ID: userID,
	}
	ctx := DBClient.Model(&model.User{}).Preload("Chats").Find(&user)

	return user.Chats, ctx.Error
}
