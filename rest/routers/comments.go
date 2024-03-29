package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/rest/controllers"
)

func CommentsRouter(api *fiber.Router) {
	co := (*api).Group("/comments")
	co.Get("/", controllers.CommentList)
	co.Get("/:id", controllers.CommentRead)
	co.Post("/", controllers.CommentCreate)
	co.Put("/:id", controllers.CommentUpdate)
	co.Patch("/:id", controllers.CommentPatch)
	co.Delete("/:id", controllers.CommentDelete)
}
