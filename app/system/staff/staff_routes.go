package staff

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(route fiber.Router) {
	handler := NewHandler()
	r := route.Group("staffs")

	r.Post("/", handler.Create)
	r.Put("/", handler.Update)
	r.Delete("/", handler.Delete)
	r.Get("/", handler.QueryPage)
	// r.Post("/roles", handler.AssignRole)
	//r.Get("/info", controller.FindInfo)
	//staff.Get("/info", handler.Info)
	//staff.Get("/role/:id", handler.FindRole)
	//staff.Post("/assign", handler.Assign)
}
