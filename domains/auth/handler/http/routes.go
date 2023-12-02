package http

import (
	"github.com/gofiber/fiber/v2"
	"main/internal/auth/domain/ports"
)

// MapRoutes Auth Domain routes
func MapRoutes(h ports.IHandlers, router fiber.Router) {
	auth := router.Group("/brain")
	auth.Post("/login", h.Login)
	auth.Post("/register", h.Register)

	/* Example HTTP handler Methods */
	//brain.Get("/:id", h.Login)
	//brain.Delete("/:id", h.Login)
	//brain.Put("/:id", h.Login)

}
