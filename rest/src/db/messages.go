package db

import "github.com/iLopezosa/api-wars/rest/src/models"

// Updater or creates a message if the id provided within the message is found or not, respectively
func MessageUpsert(m *models.Message) error {

	ctx := DBClient.Save(m)

	return ctx.Error
}

// Gets the data of the message with the provided id
func MessageRead(id uint64) (models.MessageDTO, error) {

	var message = models.MessageDTO{}
	ctx := DBClient.Model(&models.Message{}).First(&message, id)

	return message, ctx.Error
}

// Deletes the message with the provided id
func MessageDelete(id uint64) error {

	var message = models.Message{
		ID: id,
	}
	ctx := DBClient.Delete(&message)

	return ctx.Error
}

// Gets the data of all the messages
func MessageList() ([]models.MessageDTO, error) {

	var messages []models.MessageDTO
	ctx := DBClient.Model(&models.Message{}).Find(&messages)

	return messages, ctx.Error
}

// Gets the data of the messages with the provided chat id
func MessageListByChatID(chatID uint64) ([]models.MessageDTO, error) {

	var messages []models.MessageDTO
	ctx := DBClient.Model(&models.Message{}).Where("chat_id = ?", chatID).Find(&messages)

	return messages, ctx.Error
}

// Gets the data of the messages with the provided user id
func MessageListByUserID(userID uint64) ([]models.MessageDTO, error) {

	var messages []models.MessageDTO
	ctx := DBClient.Model(&models.Message{}).Where("user_id = ?", userID).Find(&messages)

	return messages, ctx.Error
}

// Gets the data of the messages with the provided chat id and user id
func MessageListByChatIDAndUserID(chatID uint64, userID uint64) ([]models.MessageDTO, error) {

	var messages []models.MessageDTO
	ctx := DBClient.Model(&models.Message{}).Where("chat_id = ? AND user_id = ?", chatID, userID).Find(&messages)

	return messages, ctx.Error
}
