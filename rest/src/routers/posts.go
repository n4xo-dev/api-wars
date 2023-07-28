package routers

import "github.com/gofiber/fiber/v2"

func Posts(api *fiber.Router) {
	p := (*api).Group("/posts")
	p.Get("/", PostList)
	p.Get("/:id", PostRead)
	p.Post("/", PostCreate)
	p.Patch("/:id", PostUpdate)
	p.Delete("/:id", PostDelete)
	p.Get("/:id/comments", PostComments)
}
