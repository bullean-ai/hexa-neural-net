package http

import (
	"context"
	"github.com/bullean-ai/hexa-neural-net/config"
	"github.com/bullean-ai/hexa-neural-net/domains/q-learning/domain/ports"
	"github.com/bullean-ai/hexa-neural-net/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// handlerHttp Auth handlers
type handlerHttp struct {
	ctx     context.Context
	cfg     *config.Config
	service ports.IService
	logger  logger.Logger
}

var (
	Responser  fiber.Map
	StatusCode int
)

// NewHttpHandler Auth Domain HTTP handlers constructor
func NewHttpHandler(ctx context.Context, cfg *config.Config, service ports.IService, logger logger.Logger) ports.IHandlers {
	return &handlerHttp{ctx: ctx, cfg: cfg, service: service, logger: logger}
}
