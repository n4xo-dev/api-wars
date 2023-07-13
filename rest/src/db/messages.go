package db

import "github.com/iLopezosa/api-wars/rest/src/models"

// Updater or creates a message if the id provided within the message is found or not, respectively
func messageUpsert(m *models.Message) error {

	ctx := DBClient.Save(m)

	return ctx.Error
}

// Gets the data of the message with the provided id
func messageRead(id uint) (models.Message, error) {

	var message = models.Message{
		ID: id,
	}
	ctx := DBClient.First(&message)

	return message, ctx.Error
}

// Deletes the message with the provided id
func messageDelete(id uint) error {

	var message = models.Message{
		ID: id,
	}
	ctx := DBClient.Delete(&message)

	return ctx.Error
}

// Gets the data of all the messages
func messageList() ([]models.Message, error) {

	var messages []models.Message
	ctx := DBClient.Find(&messages)

	return messages, ctx.Error
}

// Gets the data of the messages with the provided chat id
func messageListByChatID(chatID uint) ([]models.Message, error) {

	var messages []models.Message
	ctx := DBClient.Where("chat_id = ?", chatID).Find(&messages)

	return messages, ctx.Error
}

// Gets the data of the messages with the provided user id
func messageListByUserID(userID uint) ([]models.Message, error) {

	var messages []models.Message
	ctx := DBClient.Where("user_id = ?", userID).Find(&messages)

	return messages, ctx.Error
}

// Gets the data of the messages with the provided chat id and user id
func messageListByChatIDAndUserID(chatID uint, userID uint) ([]models.Message, error) {

	var messages []models.Message
	ctx := DBClient.Where("chat_id = ? AND user_id = ?", chatID, userID).Find(&messages)

	return messages, ctx.Error
}
