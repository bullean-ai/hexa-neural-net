package services

import (
	"fmt"
	"github.com/bullean-ai/hexa-neural-net/config"
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/application/services/layer/neuron/synapse"
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/application/services/solver"
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/entities"
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/ports"
	"github.com/bullean-ai/hexa-neural-net/pkg/logger"
	"github.com/bullean-ai/hexa-neural-net/pkg/utils/typeconv"
	"math"
	"math/rand"
	"time"
)

var (
	err error
)

const (
	collection = "examples"
)

// serviceNeuralNet Neural Net Service
type serviceNeuralNet struct {
	cfg       *config.Config
	redisRepo ports.IRedisRepository
	logger    logger.Logger
}

// NewNeuralNetService Neural Net domain service constructor
func NewNeuralNetService(cfg *config.Config, redisRepo ports.IRedisRepository, logger logger.Logger) ports.IService {
	return &serviceNeuralNet{cfg: cfg, redisRepo: redisRepo, logger: logger}
}

func (w *serviceNeuralNet) Train() {
	rand.Seed(time.Now().UnixNano())
	//percentage := .0815 // BNBUSDT
	percentage := .001 // BTCUSDT 0.005
	pair := "BTCUSDT"
	commission := .0
	iterations := 2200
	var trainData []entities.Candle

	trainData, err = w.redisRepo.GetOpenCandlesCache(fmt.Sprintf("%s:%s", pair, "OPEN:10000"))

	//trainData = trainData[600:900]
	lineData, _, maxIndex := ChartDataRedisParser(trainData, percentage, 0)
	//maxIndex = int(math.Round(float64(maxIndex) * 1.2))
	fmt.Println("maxindex: ", maxIndex)
	n := NewNeural(&entities.Config{
		Inputs:     maxIndex,
		Layout:     []int{15, 20, 20, 20, 20, 20, 10, 5, 2}, // Sufficient for modeling (AND+OR) - with 5-6 neuron always converges
		Activation: entities.ActivationSoftmax,
		Mode:       entities.ModeMultiClass,
		Weight:     synapse.NewNormal(1e-15, 1e-15),
		Bias:       true,
	})

	trainer := NewTrainer(solver.NewAdam(0.00001, 0, 0, 1e-15), 1)
	trainer.Train(n, lineData, lineData, iterations)

	var candles []entities.Candle
	candles, err = w.redisRepo.GetOpenCandlesCache(fmt.Sprintf("%s:%s", pair, "OPEN:2500"))
	//candles, _, _, err := w.redisRepo.GetCandlesData("BTCUSDT", 5)
	/*
		for i := 0; i < maxIndex+2; {
			time.Sleep(100 * time.Millisecond)
			candle, _, _ := w.redisRepo.GetCandleData(pair)
			if len(candles) > 0 && candle.Close != candles[len(candles)-1].Close {
				candles = append(candles, candle)
				i++
				println(i)
			} else if len(candles) == 0 {
				candles = append(candles, candle)
				i++
			}
		}
		if err != nil {
			println(err.Error())
		}
	*/
	CalculateProfit := entities.CalculateProfit{}
	CalculateProfit.Profit = 100
	CalculateProfit.Iterations = iterations
	fmt.Println(candles[len(candles)-1])
	//w.PredictAll(n, trainer, candles, percentage, commission, maxIndex, &CalculateProfit)
	//avgProfit, profit, longNum, lastSignal, longPercent, errorRate, testCount = w.Predict(n, trainer, candles[len(candles)-(maxIndex+1):], percentage, 1, avgProfit, profit, longNum, lastSignal, longPercent, errorRate, testCount)

	for range time.Tick(100 * time.Millisecond) {
		/*
			candles, _, _, err := w.redisRepo.GetCandlesData("BNBUSDT", 5)
			if err != nil {
				println(err.Error())
			}
		*/

		//avgProfit, profit, longNum, lastSignal, longPercent, errorRate, testCount = w.Predict(n, trainer, candles[len(candles)-(maxIndex+1):], percentage, 1, avgProfit, profit, longNum, lastSignal, longPercent, errorRate, testCount)
		candle, _, err := w.redisRepo.GetCandleData(pair)
		if err != nil {
			println(err.Error())
		}
		if len(candles) > 0 && candle.Close != candles[len(candles)-1].Close {
			candles = candles[1:]
			candles = append(candles, candle)
			Predict(n, trainer, candles[len(candles)-(maxIndex+2):], percentage, commission, maxIndex, &CalculateProfit)
		}
	}
	/*
		fmt.Println(n.Predict(testData[len(testData)-6].Input), testData[len(testData)-6].Input)
		fmt.Println(n.Predict(testData[len(testData)-5].Input), testData[len(testData)-5].Input)
		fmt.Println(n.Predict(testData[len(testData)-4].Input), testData[len(testData)-4].Input)
		fmt.Println(n.Predict(testData[len(testData)-3].Input), testData[len(testData)-3].Input)
		fmt.Println(n.Predict(testData[len(testData)-2].Input), testData[len(testData)-2].Input)
		fmt.Println(n.Predict(testData[len(testData)-1].Input), testData[len(testData)-1].Input)
	*/
}

func Predict(
	n *Neural,
	trainer *OnlineTrainer,
	candles []entities.Candle,
	percentage float64,
	commission float64,
	maxIndex int,
	profit *entities.CalculateProfit,
) {

	testData, _, _ := ChartDataRedisParser(candles, percentage, maxIndex)
	signalRes := testData[len(testData)-1]
	Calc(n, trainer, candles[len(candles)-1], signalRes, profit, percentage, commission)
	return
}

func PredictAll(
	n *Neural,
	trainer *OnlineTrainer,
	candles []entities.Candle,
	percentage float64,
	commission float64,
	maxIndex int,
	profit *entities.CalculateProfit,
) {

	testData, _, _ := ChartDataRedisParser(candles, percentage, maxIndex)
	fmt.Println(testData)
	//signalRes := testData[len(testData)-1]
	for _, signalRes := range testData {
		Calc(n, trainer, candles[len(candles)-1], signalRes, profit, percentage, commission)
	}
	return
}

func Calc(n *Neural, trainer *OnlineTrainer, candle entities.Candle, signalRes entities.Example, profit *entities.CalculateProfit, percentage, commission float64) {
	resp := n.Predict(signalRes.Input)
	firstRes := signalRes.Response[0]
	secondRes := signalRes.Response[1]
	//fmt.Println("last signal: ", profit.LastSignal)
	if math.Round(resp[0]) == 1 && profit.LastSignal == 1 {
		profit.AvgProfit = profit.AvgProfit + signalRes.Input[0]
		profit.LongPercent += signalRes.Input[0]
		profit.ShortSignalInput = append(profit.ShortSignalInput, signalRes)
		//fmt.Println(fmt.Sprintf("%f", profit.AvgProfit), profit.LongPercent)
	}
	if math.Round(resp[0]) == 1 && profit.LastSignal != 1 {
		profit.LongPercent = -commission
		profit.BuyPrice = candle
		profit.LongNum += 1
		profit.LastSignal = 1
		profit.LongSignalInput = signalRes
		profit.ShortSignalInput = append(profit.ShortSignalInput, signalRes)
	} else if math.Round(resp[1]) == 1 && profit.LastSignal != -1 {
		profit.LastSignal = -1
		profit.AvgProfit = profit.AvgProfit + signalRes.Input[0]
		profit.LongPercent += signalRes.Input[0]
		profit.Profit += profit.Profit * (profit.LongPercent / 100)
		bestLongPos := CalcBestLongPos(profit.ShortSignalInput, percentage, commission)
		if profit.LongPercent < commission && profit.TestCount > 1 {
			profit.Iterations += 1
			bestLongPos.Shuffle()
			for _, pos := range bestLongPos {
				trainer.FeedForward(n, pos)
				trainer.BackPropagate(n, pos, profit.Iterations)
			}
		}
		profit.SellPrice = candle
		profit.ShortSignalInput = []entities.Example{signalRes}

		fmt.Println("-----------------------")
		fmt.Println(fmt.Sprintf("Alım Fiyatı:%f", profit.BuyPrice.Close))
		fmt.Println(fmt.Sprintf("Satım Fiyatı: %f", profit.SellPrice.Close))
		fmt.Println(fmt.Sprintf("Long Percent: %f", profit.LongPercent))
		fmt.Println("-----------------------")

		profit.LongPercent = .0
	}
	if firstRes != math.Round(resp[0]) && secondRes != math.Round(resp[1]) {
		profit.ErrorRate += 1
	}
	profit.TestCount += 1
	fmt.Println("error rate:", profit.ErrorRate, " ", profit.TestCount, " avgProfit: ", profit.AvgProfit/float64(profit.TestCount), "profit: ", fmt.Sprintf("%.2f $", profit.Profit), " long number: ", profit.LongNum, "responses: ", math.Round(resp[0]), math.Round(resp[1]), firstRes, secondRes, "long-percent: ", profit.LongPercent)

}

func CalcBestLongPos(data []entities.Example, percentage, commission float64) (bestLong Examples) {
	bestLong = make([]entities.Example, len(data))
	buyPos := 0
	sellPos := 0
	isDone := false

	for i, _ := range bestLong {
		bestLong[i].Input = data[i].Input
		bestLong[i].Response = append(bestLong[i].Response, 0, 1)
	}

	//for i := len(data) - 1; i >= 0; i-- {

	counter := 0
	minIndex := math.MaxInt64

	for j := 0; j < len(data); j++ {
		buyPercent := .0
		counter++
		for k := j + 1; k < len(data); k++ {
			buyPercent += data[k].Input[0]

			if minIndex > k-j {
				minIndex = k - j
			}

			if buyPercent > commission {
				buyPos = j
				sellPos = k
				counter = 0
				buyPercent = 0
				j = k
				isDone = true
				break
			}

		}
		if isDone {
			break
		}
	}
	if buyPos > 0 {
		buyPos = buyPos - 1
	}
	for j := buyPos; j < sellPos; j++ {
		bestLong[j].Response[0] = 1
		bestLong[j].Response[1] = 0
	}

	if buyPos == 0 && sellPos == 0 {
		for i, _ := range bestLong {
			bestLong[i].Input = data[i].Input
			bestLong[i].Response[0] = 1
			bestLong[i].Response[1] = 0
		}
	}

	//}

	return
}

func ChartDataParser(arr []map[string]interface{}, percentage float64) (Linedata Examples, changeLine []float64, maxIndex int) {
	var longSignals, shortSignals []int
	for i := 0; i < len(arr); i++ {
		input := (typeconv.ToFloat(arr[i]["close"]) - typeconv.ToFloat(arr[i]["open"])) / typeconv.ToFloat(arr[i]["open"]) * 100
		changeLine = append(changeLine, input)
		//Linedata = append(Linedata, input)
	}
	_, longSignals, shortSignals, maxIndex = CalculateMaxPercentageDiffIndexes(changeLine, percentage, maxIndex)

	for i := maxIndex; i < len(changeLine); i++ {
		var inputExample entities.Example
		var inputs []float64
		for j := 0; j < maxIndex; j++ {
			inputs = append(inputs, changeLine[i-j])
		}
		inputExample = entities.Example{
			Input: inputs,
			Response: []float64{
				float64(longSignals[i]),
				float64(shortSignals[i]),
			},
		}
		Linedata = append(Linedata, inputExample)
	}

	return
}
func ChartDataRedisParser(arr []entities.Candle, percentage float64, maxIndex int) (Linedata Examples, changeLine []float64, maxIndexRes int) {
	var longSignals, shortSignals []int
	changeLine = CandleToChangePercent(arr)
	_, longSignals, shortSignals, maxIndexRes = CalculateMaxPercentageDiffIndexes(changeLine, percentage, maxIndex)

	if maxIndex == 0 {
		maxIndex = maxIndexRes
	}

	for i := maxIndex; i < len(changeLine); i++ {
		var inputExample entities.Example
		var inputs []float64
		for j := 0; j < maxIndex; j++ {
			inputs = append(inputs, changeLine[i-j])
		}
		inputExample = entities.Example{
			Input: inputs,
			Response: []float64{
				float64(longSignals[i]),
				float64(shortSignals[i]),
			},
		}
		Linedata = append(Linedata, inputExample)
	}

	return
}

func NeuralNetOutParser(inputs [][]float64, outs Examples, maxIndex, clusterNum int) (Linedata Examples) {

	for i := maxIndex; i < len(inputs); i++ {
		var ins []float64
		for j := i - maxIndex; j < i; j++ {
			for _, in := range inputs[j] {
				ins = append(ins, in)
			}
		}
		var inputExample entities.Example
		inputExample = entities.Example{
			Input: ins,
			Response: []float64{
				outs[i].Response[0],
				outs[i].Response[1],
			},
		}
		Linedata = append(Linedata, inputExample)
	}

	return
}

func CandleToChangePercent(arr []entities.Candle) (changeLine []float64) {
	for i := 1; i < len(arr); i++ {
		input := (arr[i].Close - arr[i-1].Close) / arr[i-1].Close * 100
		changeLine = append(changeLine, input)
		//Linedata = append(Linedata, input)
	}
	return
}

func CalculateMaxPercentageDiffIndexes(data []float64, percentage float64, maxI int) (signalPoints, longSignals, shortSignals []int, maxIndex int) {
	indexCount := 0
	maxIndex = maxI
	signalPoints = make([]int, len(data))
	buyPos := len(data) - 2
	sellPos := len(data) - 1
	isDone := false
	isDoneCount := 0
	for i := len(data) - 3; i >= 0; i-- {

		counter := 0
		minIndex := math.MaxInt64

		for j := i; j < buyPos-1; j++ {
			buyPercent := .0
			counter++
			isStarted := false
			for k := j + 1; k < buyPos; k++ {
				if isStarted {
					buyPercent += data[k]
				}
				isStarted = true

				if minIndex > k-j {
					minIndex = k - j
				}
				if buyPercent > percentage {
					if j > 0 {
						buyPos = j - 1
					} else {
						buyPos = j
					}
					if isDoneCount == 0 {
						sellPos = k
					}
					counter = 0
					buyPercent = 0
					j = k
					isDone = true
					isStarted = false
					isDoneCount += 1
					break
				}

			}

			if isDone && isDoneCount == 3 {
				signalPoints[buyPos] = 1
				signalPoints[sellPos] = -1
				isDone = false
				isDoneCount = 0
				break
			}
		}
		indexCount += 1
	}

	lastSignal := 1

	for i := 0; i < len(signalPoints); i++ {
		if lastSignal != signalPoints[i] && signalPoints[i] != 0 {
			lastSignal = signalPoints[i]
		}
		signalPoints[i] = lastSignal
	}

	lastSignal = -1
	longSignals = append(longSignals, 0)
	shortSignals = append(shortSignals, 0)
	for i := 1; i < len(signalPoints); i++ {
		if lastSignal != signalPoints[i] {
			lastSignal = signalPoints[i]
		} else if lastSignal == signalPoints[i] {
			signalPoints[i] = 0
		}
		if lastSignal == 1 {
			longSignals = append(longSignals, 1)
			shortSignals = append(shortSignals, 0)
		} else if lastSignal == -1 {
			longSignals = append(longSignals, 0)
			shortSignals = append(shortSignals, 1)
		}
	}
	return
}

/*
	func ReScaleData(list Examples) (Examples, float64, float64) {
		lastList := []float64{}
		lastList1 := []float64{}
		lastList2 := []float64{}
		lastList3 := []float64{}
		lastList4 := []float64{}

		for _, el := range list {
			lastList = append(lastList, el.Input[0])
			lastList1 = append(lastList1, el.Input[1])
			lastList2 = append(lastList1, el.Input[2])
			lastList3 = append(lastList1, el.Input[3])
			lastList4 = append(lastList1, el.Input[4])
		}
		range_value, _ := utils.MaxValue(lastList) - utils.MinValue(lastList)
		range_value = range_value + range_value/50
		min_value, _ := utils.MinValue(lastList) - range_value/100

		for i, _ := range lastList {
			list[i].Input[0] = (lastList[i] - min_value) / range_value
			list[i].Input[1] = (lastList1[i] - min_value) / range_value
			list[i].Input[2] = (lastList2[i] - min_value) / range_value
			list[i].Input[3] = (lastList3[i] - min_value) / range_value
			list[i].Input[4] = (lastList4[i] - min_value) / range_value
		}
		return list, range_value, min_value
	}
*/
