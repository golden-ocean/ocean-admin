package role

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(route fiber.Router) {
	handler := NewHandler()
	r := route.Group("roles")

	r.Post("/", handler.Create)
	r.Put("/", handler.Update)
	r.Delete("/", handler.Delete)
	r.Get("/", handler.QueryPage)
	r.Get("/all", handler.QueryAll)
	r.Get("/menus/:role_id", handler.QueryMenus)
	r.Put("/menus", handler.GrantMenus)
}
