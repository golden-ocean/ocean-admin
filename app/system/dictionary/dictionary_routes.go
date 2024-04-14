package dictionary

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(route fiber.Router) {
	handler := NewHandler()
	r := route.Group("dictionaries")

	r.Post("/", handler.Create)
	r.Put("/", handler.Update)
	r.Delete("/", handler.Delete)
	r.Get("/", handler.QueryPage)
	r.Get("/code/:code", handler.QueryByCode)
}
