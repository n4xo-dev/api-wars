package controllers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/rest/src/db"
	"gorm.io/gorm"
)

func ChatList(c *fiber.Ctx) error {
	eager := c.Query("eager", "false") == "true"

	chats, err := db.ChatList(eager)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	} else {
		return c.JSON(chats)
	}
}
func ChatRead(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	if id < 1 {
		return c.Status(400).JSON(fiber.Map{
			"error": "id must be greater than 0",
		})
	}

	eager := c.Query("eager", "false") == "true"

	chat, err := db.ChatRead(id, eager)

	if err == nil {
		return c.JSON(chat)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "chat not found",
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
}
func ChatCreate(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func ChatUpdate(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func ChatPatch(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func ChatDelete(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func ChatMessages(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	if id < 1 {
		return c.Status(400).JSON(fiber.Map{
			"error": "id must be greater than 0",
		})
	}

	messages, err := db.MessageListByChatID(id)

	if err == nil {
		return c.JSON(messages)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "chat not found",
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
}
func ChatUserMessages(c *fiber.Ctx) error {
	chatId, err := strconv.ParseUint(c.Params("chatId"), 10, 64)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "chatId is required",
		})
	}

	if chatId < 1 {
		return c.Status(400).JSON(fiber.Map{
			"error": "chatId must be greater than 0",
		})
	}

	userId, err := strconv.ParseUint(c.Params("userId"), 10, 64)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "userId is required",
		})
	}

	if userId < 1 {
		return c.Status(400).JSON(fiber.Map{
			"error": "userId must be greater than 0",
		})
	}

	messages, err := db.MessageListByChatID(chatId)

	if err == nil {
		return c.JSON(messages)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "chat or user not found",
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
}
