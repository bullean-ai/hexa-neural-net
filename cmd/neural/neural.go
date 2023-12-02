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
		Layout:     []int{15, 30, 60, 30, 15, 2}, // Sufficient for modeling (AND+OR) - with 5-6 neuron always converges
		Activation: entities.ActivationTanh,
		Mode:       entities.ModeMultiClass,
		Weight:     synapse.NewNormal(1e-15, 0),
		Bias:       true,
	})

	trainer := services.NewTrainer(solver.NewAdam(0.0001, 0, 0, 1e-15), 1)
	return neuralNetService, trainer
}
