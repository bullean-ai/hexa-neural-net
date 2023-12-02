package http

import (
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/ports"
	"github.com/gofiber/fiber/v2"
)

// MapRoutes Auth Domain routes
func MapRoutes(h ports.IHandlers, router fiber.Router) {
	//neuralNet := router.Group("/brain")

	/* Example HTTP handler Methods */
	//brain.Get("/:id", h.Login)
	//brain.Delete("/:id", h.Login)
	//brain.Put("/:id", h.Login)

}
