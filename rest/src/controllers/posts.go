package controllers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/rest/src/db"
	"gorm.io/gorm"
)

func PostList(c *fiber.Ctx) error {
	posts, err := db.PostList()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	} else {
		return c.JSON(posts)
	}
}
func PostRead(c *fiber.Ctx) error {
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

	p, err := db.PostRead(id)

	if err == nil {
		return c.JSON(p)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "post not found",
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
}
func PostCreate(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func PostUpdate(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func PostPatch(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func PostDelete(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func PostComments(c *fiber.Ctx) error {
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

	comments, err := db.CommentListByPostID(id)

	if err == nil {
		return c.JSON(comments)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "post not found",
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
}
