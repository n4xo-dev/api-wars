package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n4xo-dev/api-wars/lib/db"
	"github.com/n4xo-dev/api-wars/lib/models"
	"github.com/redis/go-redis/v9"
)

func RedisPing(c *fiber.Ctx) error {
	val, err := db.RedisPing()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).SendString(val)
}

func RedisSet(c *fiber.Ctx) error {
	var rr = new(models.RedisRecord)
	if err := c.BodyParser(rr); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := db.RedisSet(rr.Key, rr.Value)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).SendString("OK")
}

func RedisGet(c *fiber.Ctx) error {
	key := c.Params("key")

	val, err := db.RedisGet(key)
	if err == redis.Nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "key not found",
		})
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).SendString(val)
}
