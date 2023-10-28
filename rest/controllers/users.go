package controllers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/lib/db"
	"github.com/iLopezosa/api-wars/lib/models"
	"gorm.io/gorm"
)

func UserList(c *fiber.Ctx) error {
	users, err := db.UserList()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(users)
}

func UserRead(c *fiber.Ctx) error {
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

	u, err := db.UserRead(id)

	if err == nil {
		return c.JSON(u)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func UserCreate(c *fiber.Ctx) error {
	uDTO := new(models.WriteUserDTO)

	if err := c.BodyParser(uDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	u := uDTO.ToUser()

	if err := db.UserUpsert(&u); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(u.ToReadUserDTO())
}

func UserDelete(c *fiber.Ctx) error {
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

	if err := db.UserDelete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(204)
}

func UserUpdate(c *fiber.Ctx) error {
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

	uDTO := new(models.WriteUserDTO)

	if err := c.BodyParser(uDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	u := uDTO.ToUser()
	u.ID = id

	if err = db.UserUpsert(&u); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(u.ToReadUserDTO())
}

func UserPatch(c *fiber.Ctx) error {
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

	uDTO := new(models.WriteUserDTO)

	if err := c.BodyParser(uDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	u := uDTO.ToUser()
	u.ID = id

	if err = db.UserPatch(&u); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(u.ToReadUserDTO())
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

	if err == nil {
		return c.JSON(posts)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
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

	if err == nil {
		return c.JSON(comments)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
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

	if err == nil {
		return c.JSON(messages)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func UserChatMessages(c *fiber.Ctx) error {

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

	messages, err := db.MessageListByChatIDAndUserID(chatId, userId)

	if err == nil {
		return c.JSON(messages)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "user or chat not found",
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
}
