package comment

import (
	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router) {
	router_group := app.Group("/comment")
	router_group.Get("/", Index)
	router_group.Post("/get", Get)
	router_group.Post("/add", Add)
	router_group.Post("/delete/id", DelById)
	router_group.Post("/delete/page", DelByPage)
}
