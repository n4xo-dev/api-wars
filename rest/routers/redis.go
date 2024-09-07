package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n4xo-dev/api-wars/rest/controllers"
)

func RedisRouter(api *fiber.Router) {
	r := (*api).Group("/redis")
	r.Get("/", controllers.RedisPing)
	r.Post("/", controllers.RedisSet)
	r.Get("/:key", controllers.RedisGet)
}
