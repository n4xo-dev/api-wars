package controllers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/rest/src/db"
	"gorm.io/gorm"
)

func UserList(c *fiber.Ctx) error {
	users, err := db.UserList()
	if err != nil {
		return c.SendStatus(500)
	} else {
		return c.JSON(users)
	}
}
func UserRead(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	} else if id < 1 {
		return c.Status(400).JSON(fiber.Map{
			"error": "id must be greater than 0",
		})
	}

	u, err := db.UserRead(id)

	if err == nil {
		return c.JSON(u)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.SendStatus(500)
}
func UserCreate(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func UserDelete(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func UserPatch(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func UserPosts(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	} else if id < 1 {
		return c.Status(400).JSON(fiber.Map{
			"error": "id must be greater than 0",
		})
	}

	posts, err := db.PostListByUserID(id)

	if err != nil {
		return c.SendStatus(500)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.JSON(posts)
}
func UserComments(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	} else if id < 1 {
		return c.Status(400).JSON(fiber.Map{
			"error": "id must be greater than 0",
		})
	}

	comments, err := db.CommentListByUserID(id)

	if err != nil {
		return c.SendStatus(500)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.JSON(comments)
}
func UserMessages(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	} else if id < 1 {
		return c.Status(400).JSON(fiber.Map{
			"error": "id must be greater than 0",
		})
	}

	messages, err := db.MessageListByUserID(id)

	if err != nil {
		return c.SendStatus(500)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.JSON(messages)
}
func UserChatMessages(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	} else if userId < 1 {
		return c.Status(400).JSON(fiber.Map{
			"error": "id must be greater than 0",
		})
	}

	chatId, err := strconv.ParseUint(c.Params("chatId"), 10, 64)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "chatId is required",
		})
	} else if chatId < 1 {
		return c.Status(400).JSON(fiber.Map{
			"error": "chatId must be greater than 0",
		})
	}

	messages, err := db.MessageListByChatIDAndUserID(chatId, userId)

	if err != nil {
		return c.SendStatus(500)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "user or chat not found",
		})
	}

	return c.JSON(messages)
}
