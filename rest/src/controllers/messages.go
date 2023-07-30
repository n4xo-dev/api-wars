package controllers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/rest/src/db"
	"gorm.io/gorm"
)

func MessageList(c *fiber.Ctx) error {
	messages, err := db.MessageList()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	} else {
		return c.JSON(messages)
	}
}
func MessageRead(c *fiber.Ctx) error {
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

	m, err := db.MessageRead(id)

	if err == nil {
		return c.JSON(m)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "message not found",
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
}
func MessageCreate(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func MessageUpdate(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func MessageDelete(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
