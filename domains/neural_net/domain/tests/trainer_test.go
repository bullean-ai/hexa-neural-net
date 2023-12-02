package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"main/internal/neural_net/application/services"
	"main/internal/neural_net/application/services/layer/neuron/synapse"
	"main/internal/neural_net/application/services/solver"
	"main/internal/neural_net/domain/entities"
	"main/internal/neural_net/domain/utils"
	"math"
	"math/rand"
	"testing"
)

func Test_BoundedRegression(t *testing.T) {
	rand.Seed(0)

	funcs := []func(float64) float64{
		math.Sin,
		func(x float64) float64 { return math.Pow(x, 2) },
		math.Sqrt,
	}

	for _, f := range funcs {

		data := services.Examples{}
		for i := 0.0; i < 1; i += 0.01 {
			data = append(data, services.Example{Input: []float64{i}, Response: []float64{f(i)}})
		}
		n := services.NewNeural(&entities.Config{
			Inputs:     1,
			Layout:     []int{4, 4, 1},
			Activation: entities.ActivationTanh,
			Mode:       entities.ModeRegression,
			Weight:     synapse.NewUniform(0.5, 0),
			Bias:       true,
		})

		trainer := services.NewTrainer(solver.NewSGD(0.25, 0.5, 0, false), 0)
		trainer.Train(n, data, nil, 5000)

		tests := []float64{0.0, 0.1, 0.25, 0.5, 0.75, 0.9}
		for _, x := range tests {
			assert.InEpsilon(t, f(x)+1, n.Predict([]float64{x})[0]+1, 0.1)
		}
	}
}

func Test_RegressionLinearOuts(t *testing.T) {
	rand.Seed(0)
	squares := services.Examples{}
	for i := 0.0; i < 100.0; i++ {
		squares = append(squares, services.Example{Input: []float64{i}, Response: []float64{math.Sqrt(i)}})
	}
	squares.Shuffle()
	n := services.NewNeural(&entities.Config{
		Inputs:     1,
		Layout:     []int{3, 3, 1},
		Activation: entities.ActivationReLU,
		Mode:       entities.ModeRegression,
		Weight:     synapse.NewNormal(0.5, 0.5),
		Bias:       true,
	})
	trainer := services.NewBatchTrainer(solver.NewAdam(0.01, 0, 0, 0), 0, 25, 2)
	trainer.Train(n, squares, nil, 25000)
	fmt.Println(fmt.Sprintf("%v", n.Predict([]float64{0.2})))

	for i := 0; i < 100; i++ {
		x := float64(rand.Intn(99) + 1)
		assert.InEpsilon(t, math.Sqrt(x)+1, n.Predict([]float64{x})[0]+1, 0.1)
	}
}

func Test_Training(t *testing.T) {
	rand.Seed(0)

	data := services.Examples{
		services.Example{[]float64{0}, []float64{0}},
		services.Example{[]float64{0}, []float64{0}},
		services.Example{[]float64{0}, []float64{0}},
		services.Example{[]float64{5}, []float64{1}},
		services.Example{[]float64{5}, []float64{1}},
	}

	n := services.NewNeural(&entities.Config{
		Inputs:     1,
		Layout:     []int{5, 1},
		Activation: entities.ActivationSigmoid,
		Weight:     synapse.NewUniform(0.5, 0),
		Bias:       true,
	})

	trainer := services.NewTrainer(solver.NewSGD(0.5, 0.1, 0, false), 0)
	trainer.Train(n, data, nil, 1000)

	v := n.Predict([]float64{0})
	assert.InEpsilon(t, 1, 1+v[0], 0.1)
	v = n.Predict([]float64{5})
	assert.InEpsilon(t, 1.0, v[0], 0.1)
}

var data = services.Examples{
	{[]float64{2.7810836, 2.550537003}, []float64{0}},
	{[]float64{1.465489372, 2.362125076}, []float64{0}},
	{[]float64{3.396561688, 4.400293529}, []float64{0}},
	{[]float64{1.38807019, 1.850220317}, []float64{0}},
	{[]float64{3.06407232, 3.005305973}, []float64{0}},
	{[]float64{7.627531214, 2.759262235}, []float64{1}},
	{[]float64{5.332441248, 2.088626775}, []float64{1}},
	{[]float64{6.922596716, 1.77106367}, []float64{1}},
	{[]float64{8.675418651, -0.242068655}, []float64{1}},
	{[]float64{7.673756466, 3.508563011}, []float64{1}},
}

func Test_Prediction(t *testing.T) {
	rand.Seed(0)

	n := services.NewNeural(&entities.Config{
		Inputs:     2,
		Layout:     []int{2, 2, 1},
		Activation: entities.ActivationSigmoid,
		Weight:     synapse.NewUniform(0.5, 0),
		Bias:       true,
	})
	trainer := services.NewTrainer(solver.NewSGD(0.5, 0.1, 0, false), 0)

	trainer.Train(n, data, nil, 5000)

	for _, d := range data {
		assert.InEpsilon(t, n.Predict(d.Input)[0]+1, d.Response[0]+1, 0.1)
	}
}

func Test_CrossVal(t *testing.T) {
	n := services.NewNeural(&entities.Config{
		Inputs:     2,
		Layout:     []int{1, 1},
		Activation: entities.ActivationTanh,
		Loss:       entities.LossMeanSquared,
		Weight:     synapse.NewUniform(0.5, 0),
		Bias:       true,
	})

	trainer := services.NewTrainer(solver.NewSGD(0.5, 0.1, 0, false), 0)
	trainer.Train(n, data, data, 1000)

	for _, d := range data {
		assert.InEpsilon(t, n.Predict(d.Input)[0]+1, d.Response[0]+1, 0.1)
		assert.InEpsilon(t, 1, services.CrossValidate(n, data)+1, 0.01)
	}
}

func Test_MultiClass(t *testing.T) {
	var data = services.Examples{
		{[]float64{2.7810836, 2.550537003}, []float64{1, 0}},
		{[]float64{1.465489372, 2.362125076}, []float64{1, 0}},
		{[]float64{3.396561688, 4.400293529}, []float64{1, 0}},
		{[]float64{1.38807019, 1.850220317}, []float64{1, 0}},
		{[]float64{3.06407232, 3.005305973}, []float64{1, 0}},
		{[]float64{7.627531214, 2.759262235}, []float64{0, 1}},
		{[]float64{5.332441248, 2.088626775}, []float64{0, 1}},
		{[]float64{6.922596716, 1.77106367}, []float64{0, 1}},
		{[]float64{8.675418651, -0.242068655}, []float64{0, 1}},
		{[]float64{7.673756466, 3.508563011}, []float64{0, 1}},
	}

	n := services.NewNeural(&entities.Config{
		Inputs:     2,
		Layout:     []int{3, 2},
		Activation: entities.ActivationReLU,
		Mode:       entities.ModeMultiClass,
		Loss:       entities.LossMeanSquared,
		Weight:     synapse.NewUniform(0.1, 0),
		Bias:       true,
	})

	trainer := services.NewTrainer(solver.NewSGD(0.01, 0.1, 0, false), 0)
	trainer.Train(n, data, data, 1000)

	for _, d := range data {
		est := n.Predict(d.Input)
		assert.InEpsilon(t, 1.0, utils.Sum(est), 0.00001)
		if d.Response[0] == 1.0 {
			assert.InEpsilon(t, n.Predict(d.Input)[0]+1, d.Response[0]+1, 0.1)
		} else {
			assert.InEpsilon(t, n.Predict(d.Input)[1]+1, d.Response[1]+1, 0.1)
		}
		assert.InEpsilon(t, 1, services.CrossValidate(n, data)+1, 0.01)
	}

}

func Test_or(t *testing.T) {
	rand.Seed(0)
	n := services.NewNeural(&entities.Config{
		Inputs:     2,
		Layout:     []int{1, 1},
		Activation: entities.ActivationTanh,
		Mode:       entities.ModeBinary,
		Weight:     synapse.NewUniform(0.5, 0),
		Bias:       true,
	})
	permutations := services.Examples{
		{[]float64{0, 0}, []float64{0}},
		{[]float64{1, 0}, []float64{1}},
		{[]float64{0, 1}, []float64{1}},
		{[]float64{1, 1}, []float64{1}},
	}

	trainer := services.NewTrainer(solver.NewSGD(0.5, 0, 0, false), 10)

	trainer.Train(n, permutations, permutations, 25)

	for _, perm := range permutations {
		assert.Equal(t, utils.Round(n.Predict(perm.Input)[0]), perm.Response[0])
	}
	fmt.Println(n.Predict([]float64{0, 1}))

}

func Test_xor(t *testing.T) {
	rand.Seed(0)
	n := services.NewNeural(&entities.Config{
		Inputs:     2,
		Layout:     []int{5, 1}, // Sufficient for modeling (AND+OR) - with 5-6 neuron always converges
		Activation: entities.ActivationSigmoid,
		Mode:       entities.ModeBinary,
		Weight:     synapse.NewUniform(.25, 0),
		Bias:       true,
	})
	permutations := services.Examples{
		{[]float64{0, 0}, []float64{0}},
		{[]float64{1, 0}, []float64{1}},
		{[]float64{0, 1}, []float64{1}},
		{[]float64{1, 1}, []float64{0}},
	}

	trainer := services.NewTrainer(solver.NewSGD(1, 0.1, 1e-6, false), 50)
	trainer.Train(n, permutations, permutations, 1000)

	for _, perm := range permutations {
		assert.InEpsilon(t, n.Predict(perm.Input)[0]+1, perm.Response[0]+1, 0.2)
	}
	fmt.Println(n.Predict([]float64{0, 0})[0])

}

func printResult(ideal, actual []float64) {
	fmt.Printf("want: %+v have: %+v\n", ideal, actual)
}
