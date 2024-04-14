package dictionary_item

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(route fiber.Router) {
	handler := NewHandler()
	dictionaryRoutes := route.Group("dictionary")
	r := dictionaryRoutes.Group("items")

	r.Get("/", handler.QueryPage)
	r.Post("/", handler.Create)
	r.Put("/", handler.Update)
	r.Delete("/", handler.Delete)

}
