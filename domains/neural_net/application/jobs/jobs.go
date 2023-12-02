package jobs

import (
	"context"
	"github.com/bullean-ai/hexa-neural-net/config"
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/ports"
	"github.com/bullean-ai/hexa-neural-net/pkg/logger"
)

// jobRunner Indicator Worker struct
type jobRunner struct {
	cfg    *config.Config
	logger logger.Logger
	srv    ports.IService
}

// NewJobRunner Indicator worker constructor
func NewJobRunner(cfg *config.Config, logger logger.Logger, srv ports.IService) ports.IJobs {
	return &jobRunner{
		cfg:    cfg,
		logger: logger,
		srv:    srv,
	}
}

func (w *jobRunner) TrainNeuralNet(ctx context.Context) {
	w.logger.Info("Training is started")
	w.srv.Train()
}
