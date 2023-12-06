package neural

import (
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/application/services"
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/application/services/layer/neuron/synapse"
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/application/services/solver"
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/entities"
)

func Init(inputLength int) (*services.Neural, *services.OnlineTrainer) {
	// Init services
	neuralNetService := services.NewNeural(&entities.Config{
		Inputs:     inputLength,
		Layout:     []int{15, 20, 20, 20, 20, 20, 10, 5, 2}, // Sufficient for modeling (AND+OR) - with 5-6 neuron always converges
		Activation: entities.ActivationSoftmax,
		Mode:       entities.ModeMultiClass,
		Weight:     synapse.NewNormal(1e-15, 1e-15),
		Bias:       true,
	})

	trainer := services.NewTrainer(solver.NewAdam(0.00001, 0, 0, 1e-15), 1)
	return neuralNetService, trainer
}
