package routers

import "github.com/gofiber/fiber/v2"

func Messages(api *fiber.Router) {
	m := (*api).Group("/messages")
	m.Get("/", MessageList)
	m.Get("/:id", MessageRead)
	m.Post("/", MessageCreate)
	m.Patch("/:id", MessageUpdate)
	m.Delete("/:id", MessageDelete)
}
