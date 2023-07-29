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
	return c.SendStatus(501)
}
func UserComments(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func UserMessages(c *fiber.Ctx) error {
	return c.SendStatus(501)
}
func UserChatMessages(c *fiber.Ctx) error {
	return c.SendStatus(501)
}