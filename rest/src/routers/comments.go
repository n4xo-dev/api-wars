package routers

import "github.com/gofiber/fiber/v2"

func CommentsRouter(api *fiber.Router) {
	co := (*api).Group("/comments")
	co.Get("/", CommentList)
	co.Get("/:id", CommentRead)
	co.Post("/", CommentCreate)
	co.Patch("/:id", CommentUpdate)
	co.Delete("/:id", CommentDelete)
}
