package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golden-ocean/ocean-admin/pkg/middlewares"
)

func InitRoutes(route fiber.Router) {
	handler := NewHandler()
	r := route.Group("auth")

	r.Post("/login", handler.Login)
	// r.Post("/logout", handler.Logout)
	r.Get("/info", middlewares.JWTProtected(), handler.QueryInfo)
	r.Post("/refresh", middlewares.JWTProtected(), handler.Refresh)
	r.Get("/test", handler.Test)

}
