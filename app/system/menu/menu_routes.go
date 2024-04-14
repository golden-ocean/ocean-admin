package menu

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(route fiber.Router) {
	handler := NewHandler()
	r := route.Group("menus")

	// r.Get("/tree", handler.QueryTreeByStaffID)
	r.Post("/", handler.Create)
	r.Put("/", handler.Update)
	r.Delete("/", handler.Delete)
	r.Get("/tree", handler.QueryTree)
}
