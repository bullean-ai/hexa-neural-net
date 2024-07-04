package entities

import (
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/application/services/layer/neuron/synapse"
	"time"
)

type GetReq struct {
	Src         int8   `json:"src" validate:"required,min=1,max=5"`
	UserType    int8   `json:"user_type" validate:"required,min=1,max=5"`
	UserTitle   string `json:"user_title" validate:"required"`
	CompanyName string `json:"company_name,omitempty"`
	UserName    string `json:"user_name" validate:"required,min=10,max=50"`
	UserPass    string `json:"user_pass,omitempty"`
	UserPhone   string `json:"user_phone" validate:"required,min=10,max=20"`
	VerifyCode  int64  `json:"verify_code,omitempty"`
}

type HandlerResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Mode denotes inference mode
type Mode int

const (
	// ModeDefault is unspecified mode
	ModeDefault Mode = 0
	// ModeMultiClass is for one-hot encoded classification, applies softmax output layer
	ModeMultiClass Mode = 1
	// ModeRegression is regression, applies linear output layer
	ModeRegression Mode = 2
	// ModeBinary is binary classification, applies sigmoid output layer
	ModeBinary Mode = 3
	// ModeMultiLabel is for multilabel classification, applies sigmoid output layer
	ModeMultiLabel Mode = 4
)

// ActivationType is represents a neuron activation function
type ActivationType int

const (
	// ActivationNone is no activation
	ActivationNone ActivationType = 0
	// ActivationSigmoid is a sigmoid activation
	ActivationSigmoid ActivationType = 1
	// ActivationTanh is hyperbolic activation
	ActivationTanh ActivationType = 2
	// ActivationReLU is rectified linear unit activation
	ActivationReLU ActivationType = 3
	// ActivationLinear is linear activation
	ActivationLinear ActivationType = 4
	// ActivationSoftmax is a softmax activation (per layer)
	ActivationSoftmax ActivationType = 5
)

// Config defines the network topology, activations, losses etc
type Config struct {
	// Number of inputs
	Inputs int
	// Defines topology:
	// For instance, [5 3 3] signifies a network with two hidden layers
	// containing 5 and 3 nodes respectively, followed an output layer
	// containing 3 nodes.
	Layout []int
	// Activation functions: {ActivationTanh, ActivationReLU, ActivationSigmoid}
	Activation ActivationType
	// Solver modes: {ModeRegression, ModeBinary, ModeMultiClass, ModeMultiLabel}
	Mode Mode
	// Initializer for weights: {NewNormal(σ, μ), NewUniform(σ, μ)}
	Weight synapse.WeightInitializer
	// Loss functions: {LossCrossEntropy, LossBinaryCrossEntropy, LossMeanSquared}
	Loss LossType
	// Apply bias nodes
	Bias bool
}

// LossType represents a loss function
type LossType int

func (l LossType) String() string {
	switch l {
	case LossCrossEntropy:
		return "CE"
	case LossBinaryCrossEntropy:
		return "BinCE"
	case LossMeanSquared:
		return "MSE"
	}
	return "N/A"
}

const (
	// LossNone signifies unspecified loss
	LossNone LossType = 0
	// LossCrossEntropy is cross entropy loss
	LossCrossEntropy LossType = 1
	// LossBinaryCrossEntropy is the special case of binary cross entropy loss
	LossBinaryCrossEntropy LossType = 2
	// LossMeanSquared is MSE
	LossMeanSquared LossType = 3
)

const EXCHANGE_TYPE = "BINANCE_SPOT"

type Candle struct {
	Date             *time.Time
	Open             float64
	High             float64
	Low              float64
	Close            float64
	Volume           float64
	QuoteAssetVolume float64
	TakerBaseVolume  float64
	TakerQuoteVolume float64
}

type TickCandle struct {
	Price float64 `json:"price"`
}

type DepthData struct {
	Symbol string       `msg:"Symbol"`
	Bids   []DepthPrice `msg:"Bids"`
	Asks   []DepthPrice `msg:"Asks"`
}

type StatisticsModel struct {
	Pair  string `json:"pair" validate:"required"`
	Short int    `json:"short" validate:"required"`
	Long  int    `json:"long" validate:"required"`
}

type DepthPrice struct {
	// Fiyat
	Price float64
	// Alım veya satımlar
	Sale float64
}

// Example is an input-target pair
type Example struct {
	Input    []float64
	Response []float64
}
type CalculateProfit struct {
	LastSignal       int64      `json:"last_signal"`
	AvgProfit        float64    `json:"avg_profit"`
	LongPercent      float64    `json:"long_percent"`
	LongNum          int64      `json:"long_num"`
	LongSignalInput  Example    `json:"long_signal_input"`
	ShortSignalInput []Example  `json:"short_signal_input"`
	Profit           float64    `json:"profit"`
	ErrorRate        int64      `json:"error_rate"`
	TestCount        int64      `json:"test_count"`
	BuyPrice         TickCandle `json:"buy_price"`
	SellPrice        TickCandle `json:"sell_price"`
	Iterations       int        `json:"iterations"`
}
