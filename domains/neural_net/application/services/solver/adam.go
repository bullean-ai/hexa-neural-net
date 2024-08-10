package solver

import (
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/utils"
	"math"
)

// Adam is an Adam solver
type Adam struct {
	lr      *float64
	beta    *float64
	beta2   *float64
	epsilon *float64

	v, m []*float64
}

// NewAdam returns a new Adam solver
func NewAdam(lr, beta, beta2, epsilon float64) *Adam {
	lr = utils.Fparam(lr, 0.001)
	beta = utils.Fparam(beta, 0.9)
	beta2 = utils.Fparam(beta2, 0.999)
	epsilon = utils.Fparam(epsilon, 1e-8)

	return &Adam{
		lr:      &lr,
		beta:    &beta,
		beta2:   &beta2,
		epsilon: &epsilon,
	}
}

// Init initializes vectors using number of weights in network
func (o *Adam) Init(size int) {
	o.v, o.m = make([]*float64, size), make([]*float64, size)
}

// Update returns the update for a given weight
func (o *Adam) Update(value, gradient float64, t, idx int) float64 {
	lrt := *o.lr * (math.Sqrt(1.0 - math.Pow(*o.beta2, float64(t)))) /
		(1.0 - math.Pow(*o.beta, float64(t)))
	oM := (*o.beta)*(*o.m[idx]) + (1.0-(*o.beta))*gradient
	o.m[idx] = &oM
	oV := (*o.beta2)*(*o.v[idx]) + (1.0-(*o.beta2))*math.Pow(gradient, 2.0)
	o.v[idx] = &oV
	retVal := -lrt * ((*o.m[idx]) / (math.Sqrt(*o.v[idx]) + (*o.epsilon)))
	return retVal
}
