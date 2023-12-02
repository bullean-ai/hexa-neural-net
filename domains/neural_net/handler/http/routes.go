package http

import (
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/ports"
	"github.com/gofiber/fiber/v2"
)

// MapRoutes Auth Domain routes
func MapRoutes(h ports.IHandlers, router fiber.Router) {
	//neuralNet := router.Group("/auth")

	/* Example HTTP handler Methods */
	//auth.Get("/:id", h.Login)
	//auth.Delete("/:id", h.Login)
	//auth.Put("/:id", h.Login)

}
