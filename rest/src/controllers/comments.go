package controllers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/rest/src/db"
	"gorm.io/gorm"
)

func CommentList(c *fiber.Ctx) error {
	comments, err := db.CommentList()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	} else {
		return c.JSON(comments)
	}
}
func CommentRead(c *fiber.Ctx) error {
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

	comm, err := db.CommentRead(id)

	if err == nil {
		return c.JSON(comm)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "comment not found",
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
}
func CommentCreate(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func CommentUpdate(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func CommentDelete(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
