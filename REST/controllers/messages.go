package controllers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/rest/db"
	"github.com/iLopezosa/api-wars/rest/models"
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
	msgDTO := new(models.WriteMessageDTO)

	if err := c.BodyParser(msgDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	msg := msgDTO.ToMessage()

	if err := db.MessageUpsert(&msg); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(msg.ToReadMessageDTO())
}

func MessageUpdate(c *fiber.Ctx) error {
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

	msgDTO := new(models.WriteMessageDTO)

	if err := c.BodyParser(msgDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	msg := msgDTO.ToMessage()
	msg.ID = id

	if err := db.MessageUpsert(&msg); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(msg.ToReadMessageDTO())
}

func MessagePatch(c *fiber.Ctx) error {
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

	msgDTO := new(models.WriteMessageDTO)

	if err := c.BodyParser(msgDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	msg := msgDTO.ToMessage()
	msg.ID = id

	if err := db.MessagePatch(&msg); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(msg.ToReadMessageDTO())
}

func MessageDelete(c *fiber.Ctx) error {
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

	if err := db.MessageDelete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(204)
}
