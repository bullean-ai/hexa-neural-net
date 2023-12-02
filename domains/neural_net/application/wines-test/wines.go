package wines_test

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"main/internal/neural_net/application/services"
	"main/internal/neural_net/application/services/layer/neuron/synapse"
	"main/internal/neural_net/application/services/solver"
	"main/internal/neural_net/domain/entities"
	"main/internal/neural_net/domain/utils"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func WinesTest() {

	rand.Seed(time.Now().UnixNano())

	data, err := load("./domains/neural_net/application/wines-test/wines.data")
	if err != nil {
		panic(err)
	}

	for i := range data {
		utils.Standardize(data[i].Input)
	}
	data.Shuffle()

	fmt.Printf("have %v entries\n", data[0])

	neural := services.NewNeural(&entities.Config{
		Inputs:     len(data[0].Input),
		Layout:     []int{8, 3},
		Activation: entities.ActivationTanh,
		Mode:       entities.ModeMultiClass,
		Weight:     synapse.NewNormal(1, 0),
		Bias:       true,
	})

	//trainer := training.NewTrainer(training.NewSGD(0.005, 0.5, 1e-6, true), 50)
	//trainer := training.NewBatchTrainer(training.NewSGD(0.005, 0.1, 0, true), 50, 300, 16)
	//trainer := training.NewTrainer(training.NewAdam(0.1, 0, 0, 0), 50)
	trainer := services.NewBatchTrainer(solver.NewAdam(0.1, 0, 0, 0), 50, len(data)/2, 12)
	//data, heldout := data.Split(0.5)
	trainer.Train(neural, data, data, 50000)
	fmt.Println(neural.Predict([]float64{-0.2656968524232809, -0.33798087994931847, -0.3411236637547984, -0.1959968915817489, 0.317324463313301, -0.34929490164904614, -0.356488384581589, -0.35662806386183254, -0.3562788656612237, -0.329739802414949, -0.353904317897083, -0.34573308000283554, 3.2715412404644044}))
}

func load(path string) (services.Examples, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(bufio.NewReader(f))

	var examples services.Examples
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		examples = append(examples, toExample(record))
	}

	return examples, nil
}

func toExample(in []string) services.Example {
	res, err := strconv.ParseFloat(in[0], 64)
	if err != nil {
		panic(err)
	}
	resEncoded := onehot(3, res)
	var features []float64
	for i := 1; i < len(in); i++ {
		res, err := strconv.ParseFloat(in[i], 64)
		if err != nil {
			panic(err)
		}
		features = append(features, res)
	}

	return services.Example{
		Response: resEncoded,
		Input:    features,
	}
}

func onehot(classes int, val float64) []float64 {
	res := make([]float64, classes)
	res[int(val)-1] = 1
	return res
}
