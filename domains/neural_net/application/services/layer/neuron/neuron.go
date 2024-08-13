package neuron

import (
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/application/services/layer/neuron/synapse"
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/entities"
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/utils"
)

// Neuron is a neural network node
type Neuron struct {
	A     entities.ActivationType
	In    []*synapse.Synapse
	Out   []*synapse.Synapse
	Value float64
	Index int
}

// NewNeuron returns a neuron with the given activation
func NewNeuron(activation entities.ActivationType) *Neuron {
	return &Neuron{
		A: activation,
	}
}

func (n *Neuron) Fire() {
	var sum float64
	for _, s := range n.In {
		sum += s.Out
	}
	n.Value = n.Activate(sum)

	nVal := n.Value
	for _, s := range n.Out {
		s.Fire(nVal)
	}
}

// Activate applies the neurons activation
func (n *Neuron) Activate(x float64) float64 {
	return utils.GetActivation(n.A).F(x)
}

// DActivate applies the derivative of the neurons activation
func (n *Neuron) DActivate(x float64) float64 {
	return utils.GetActivation(n.A).Df(x)
}
