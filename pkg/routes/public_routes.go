package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golden-ocean/ocean-admin/app/security/auth"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/")
	auth.InitRoutes(route)
}
