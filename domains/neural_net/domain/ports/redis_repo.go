package ports

import "github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/entities"

// IRedisRepository Whale Hunter Domain redis interface
type IRedisRepository interface {
	SetCache(string, any) error
	DeleteCache(string) error
	GetCandleData(string) (entities.Candle, string, error)
	GetCandlesData(string, int) ([]entities.Candle, int, string, error)
	GetDepthData(string) (entities.DepthData, string, error)
	GetOpenCandlesCache(string) ([]entities.Candle, error)
}
