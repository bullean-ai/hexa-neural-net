package utils

import (
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/application/services/layer/neuron/synapse/activation"
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/entities"
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/ports"
	"math"
	"math/rand"
	"strings"
)

// OutputActivation returns activation corresponding to prediction mode
func OutputActivation(c entities.Mode) entities.ActivationType {
	switch c {
	case entities.ModeMultiClass:
		return entities.ActivationSoftmax
	case entities.ModeRegression:
		return entities.ActivationLinear
	case entities.ModeBinary, entities.ModeMultiLabel:
		return entities.ActivationSigmoid
	}
	return entities.ActivationNone
}

// GetActivation returns the concrete activation given an ActivationType
func GetActivation(act entities.ActivationType) ports.Differentiable {
	switch act {
	case entities.ActivationSigmoid:
		return &activation.Sigmoid{}
	case entities.ActivationTanh:
		return &activation.Tanh{}
	case entities.ActivationReLU:
		return &activation.ReLU{}
	case entities.ActivationLinear:
		return &activation.Linear{}
	case entities.ActivationSoftmax:
		return &activation.Linear{}
	}
	return &activation.Linear{}
}

// Mean of xx
func Mean(xx []float64) float64 {
	var sum float64
	for _, x := range xx {
		sum += x
	}
	return sum / float64(len(xx))
}

// Variance of xx
func Variance(xx []float64) float64 {
	if len(xx) == 1 {
		return 0.0
	}
	m := Mean(xx)

	var variance float64
	for _, x := range xx {
		variance += math.Pow((x - m), 2)
	}

	return variance / float64(len(xx)-1)
}

// StandardDeviation of xx
func StandardDeviation(xx []float64) float64 {
	return math.Sqrt(Variance(xx))
}

// Standardize (z-score) shifts distribution to μ=0 σ=1
func Standardize(xx []float64) {
	m := Mean(xx)
	s := StandardDeviation(xx)

	if s == 0 {
		s = 1
	}

	for i, x := range xx {
		xx[i] = (x - m) / s
	}
}

// Normalize scales to (0,1)
func Normalize(xx []float64) {
	min, max := Min(xx), Max(xx)
	for i, x := range xx {
		xx[i] = (x - min) / (max - min)
	}
}

// Min is the smallest element
func Min(xx []float64) float64 {
	min := xx[0]
	for _, x := range xx {
		if x < min {
			min = x
		}
	}
	return min
}

// Max is the largest element
func Max(xx []float64) float64 {
	max := xx[0]
	for _, x := range xx {
		if x > max {
			max = x
		}
	}
	return max
}

// ArgMax is the index of the largest element
func ArgMax(xx []float64) int {
	max, idx := xx[0], 0
	for i, x := range xx {
		if x > max {
			max, idx = xx[i], i
		}
	}
	return idx
}

// Sgn is signum
func Sgn(x float64) float64 {
	switch {
	case x < 0:
		return -1.0
	case x > 0:
		return 1.0
	}
	return 0
}

// Sum is sum
func Sum(xx []float64) (sum float64) {
	for _, x := range xx {
		sum += x
	}
	return
}

// Softmax is the softmax function
func Softmax(xx []float64) []float64 {
	out := make([]float64, len(xx))
	var sum float64
	max := Max(xx)
	for i, x := range xx {
		out[i] = math.Exp(x - max)
		sum += out[i]
	}
	for i := range out {
		out[i] /= sum
	}
	return out
}

// Round to nearest integer
func Round(x float64) float64 {
	return math.Floor(x + .5)
}

// Dot product
func Dot(xx, yy []float64) float64 {
	var p float64
	for i := range xx {
		p += xx[i] * yy[i]
	}
	return p
}

func Fparam(val, fallback float64) float64 {
	if val == 0.0 {
		return fallback
	}
	return val
}

func Iparam(val, fallback int) int {
	if val == 0 {
		return fallback
	}
	return val
}

// CheckStringIfContains check a string if contains given param
func CheckStringIfContains(input_text string, search_text string) bool {
	CheckContains := strings.Contains(input_text, search_text)
	return CheckContains
}

// Sma Simple Moving Average.
func Sma(period int, values []float64) []float64 {
	result := make([]float64, len(values))
	sum := float64(0)

	for i, value := range values {
		count := i + 1
		sum += value

		if i >= period {
			sum -= values[i-period]
			count = period
		}

		result[i] = sum / float64(count)
	}

	return result
}

func MaxValue(values []float64) (result float64, index int) {
	result = math.MinInt64
	for i := 0; i < len(values); i++ {
		if values[i] > result {
			result = values[i]
			index = i

		}
	}
	return
}

func MinValue(values []float64) (result float64, index int) {
	result = math.MaxInt64
	for i := 0; i < len(values); i++ {
		if values[i] < result {
			result = values[i]
			index = i
		}
	}
	return
}

func GenerateSignalMap(length int) (mapping map[int64]int64, buyArr, sellArr []int64) {
	mapping = make(map[int64]int64)
	for i := 0; i < length; i++ {
		key := int64(rand.Float64() * 1000000)
		mapping[key] = 1
		buyArr = append(buyArr, key)
	}
	for i := 0; i < length; i++ {
		key := int64(rand.Float64() * 1000000)
		mapping[key] = -1
		sellArr = append(sellArr, key)

	}
	return
}

/*

func MarketDepth(depthData ent.DepthData, interval int) map[string]float64 {
	var buyPrice []float64
	var sellPrice []float64
	totalBuyAmount := 0.0
	totalBuyVolume := 0.0
	totalSellAmount := 0.0
	totalSellVolume := 0.0
	calculatedData := map[string]float64{}
	if interval > len(depthData.Bids) {
		interval = len(depthData.Bids)
	}
	for i := len(depthData.Bids) - interval; i < len(depthData.Bids); i++ {
		buyPrice = append(buyPrice, depthData.Bids[i].Price)
		sellPrice = append(sellPrice, depthData.Asks[i].Price)
		totalBuyAmount += depthData.Bids[i].Sale
		totalBuyVolume += depthData.Bids[i].Sale * depthData.Bids[i].Price
		totalSellAmount += depthData.Asks[i].Sale
		totalSellVolume += depthData.Asks[i].Sale * depthData.Asks[i].Price
	}
	averageBuyPrice := totalBuyVolume / totalBuyAmount
	averageSellPrice := totalSellVolume / totalSellAmount
	buyPercentage := totalBuyAmount / (totalBuyAmount + totalSellAmount) * 100
	sellPercentage := totalSellAmount / (totalBuyAmount + totalSellAmount) * 100
	if buyPercentage > sellPercentage {
		calculatedData["signal"] = 1
	} else if buyPercentage < sellPercentage {
		calculatedData["signal"] = -1
	} else {
		calculatedData["signal"] = 0
	}
	calculatedData["average_buy_price"] = averageBuyPrice
	calculatedData["average_sell_price"] = averageSellPrice
	calculatedData["buy_percentage"] = buyPercentage
	calculatedData["sell_percentage"] = sellPercentage
	calculatedData["buy_price"] = buyPrice[len(buyPrice)-1]
	calculatedData["sell_price"] = sellPrice[len(sellPrice)-1]
	return calculatedData
}
// Alış ve Satış Hacimlerini Ayırma Fonksiyonu
func SeperateVol(candle ent.Candle) (unit float64, volume []float64) {
	if candle.Close > candle.Open {
		unit = candle.Volume / (2*math.Abs(candle.High-candle.Close) + 2*math.Abs(candle.Open-candle.Low) + math.Abs(candle.Close-candle.Open))
	} else if candle.Open > candle.Close {
		unit = candle.Volume / (2*math.Abs(candle.High-candle.Open) + 2*math.Abs(candle.Close-candle.Low) + math.Abs(candle.Close-candle.Open))
	} else {
		unit = 0
	}
	diff := math.Abs(candle.Close-candle.Open) * unit
	raidVolume := (candle.Volume + diff) / 2
	if candle.Close > candle.Open {
		volume = append(volume, raidVolume)
		volume = append(volume, candle.Volume-raidVolume)
	} else if candle.Open > candle.Close {
		volume = append(volume, candle.Volume-raidVolume)
		volume = append(volume, raidVolume)
	} else if candle.Close == candle.Open {
		volume = append(volume, raidVolume)
		volume = append(volume, raidVolume)
	}
	return unit, volume // Alış ve Satış hacmini sırası ile diziye bastırdık
}
*/
