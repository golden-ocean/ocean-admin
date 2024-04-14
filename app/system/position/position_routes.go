package position

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(route fiber.Router) {
	handler := NewHandler()
	r := route.Group("positions")

	r.Post("/", handler.Create)
	r.Put("/", handler.Update)
	r.Delete("/", handler.Delete)
	r.Get("/", handler.QueryPage)
	r.Get("/all", handler.QueryAll)

}
