package routers

import "github.com/gofiber/fiber/v2"

func Chats(api *fiber.Router) {
	c := (*api).Group("/chats")
	c.Get("/", ChatList)
	c.Get("/:id", ChatRead) // Can load messages and participants w/ eager loading
	c.Post("/", ChatCreate)
	c.Patch("/:id", ChatUpdate)
	c.Delete("/:id", ChatDelete)
	c.Get("/:id/messages", ChatMessages)
	c.Get("/:id/user/:id2/messages", ChatUserMessages)
}
