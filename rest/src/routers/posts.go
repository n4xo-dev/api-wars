package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/rest/src/controllers"
)

func PostsRouter(api *fiber.Router) {
	p := (*api).Group("/posts")
	p.Get("/", controllers.PostList)
	p.Get("/:id", controllers.PostRead)
	p.Post("/", controllers.PostCreate)
	p.Patch("/:id", controllers.PostUpdate)
	p.Delete("/:id", controllers.PostDelete)
	p.Get("/:id/comments", controllers.PostComments)
}