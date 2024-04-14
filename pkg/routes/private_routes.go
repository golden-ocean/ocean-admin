package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golden-ocean/ocean-admin/app/system/dictionary"
	"github.com/golden-ocean/ocean-admin/app/system/dictionary_item"
	"github.com/golden-ocean/ocean-admin/app/system/menu"
	"github.com/golden-ocean/ocean-admin/app/system/organization"
	"github.com/golden-ocean/ocean-admin/app/system/position"
	"github.com/golden-ocean/ocean-admin/app/system/role"
	"github.com/golden-ocean/ocean-admin/app/system/staff"
	"github.com/golden-ocean/ocean-admin/pkg/middlewares"
)

func PrivateRoutes(a *fiber.App) {
	appRoute := a.Group("/")
	// system := appRoute.Group("/system")
	system := appRoute.Group("/system", middlewares.JWTProtected(), middlewares.CasbinProtected())
	// system := appRoute.Group("/system", middlewares.JWTProtected())

	dictionary.InitRoutes(system)
	dictionary_item.InitRoutes(system)
	staff.InitRoutes(system)
	role.InitRoutes(system)
	position.InitRoutes(system)
	organization.InitRoutes(system)
	menu.InitRoutes(system)
}
