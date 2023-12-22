package repository

import (
	"context"
	"fmt"
	ent1 "github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/entities"
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/ports"
	ent "github.com/bullean-ai/hexa-neural-net/domains/neural_net/infrastructure/adapters/binance"
	"github.com/bullean-ai/hexa-neural-net/pkg/utils/typeconv"
	"github.com/go-redis/redis/v8"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
)

// redisRepo Struct
type redisRepo struct {
	ctx         context.Context
	redisClient *redis.Client
}

// NewRedisRepo Whale Hunter Domain redis repository constructor
func NewRedisRepo(ctx context.Context, redisClient *redis.Client) ports.IRedisRepository {
	return &redisRepo{
		ctx:         ctx,
		redisClient: redisClient,
	}
}

// GetCacheFeatureKline Get Kline of features cache by key
func (n *redisRepo) GetCacheFeatureKline(key string) ([]ent.ChartData, error) {
	bytes, err := n.redisClient.Get(n.ctx, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.GetCache")
	}

	var base []ent.ChartData
	if err = json.Unmarshal(bytes, &base); err != nil {
		return nil, errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.GetCache.json.Unmarshal")
	}

	return base, nil
}

// GetCacheLongShort Get cache of long short by key
func (n *redisRepo) GetCacheLongShort(key string) ([]ent.TakerLongShortRatioData, error) {
	indicatorsBytes, err := n.redisClient.Get(n.ctx, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.GetCache")
	}

	indicatorsBase := []ent.TakerLongShortRatioData{}
	if err = json.Unmarshal(indicatorsBytes, &indicatorsBase); err != nil {
		return nil, errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.GetCache.json.Unmarshal")
	}

	return indicatorsBase, nil
}

// GetCacheDepth Get Depth cache by key
func (n *redisRepo) GetCacheDepth(key string) (base ent.DepthData, err error) {
	bytes, err := n.redisClient.Get(n.ctx, key).Bytes()
	if err != nil {
		return ent.DepthData{}, errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.GetCache")
	}

	if err = json.Unmarshal(bytes, &base); err != nil {
		return ent.DepthData{}, errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.GetCache.json.Unmarshal")
	}

	return
}

func (n *redisRepo) GetCandleData(symbol string) (ent1.Candle, string, error) {
	redisKey := fmt.Sprintf("%s:%s:%d:Open", ent1.EXCHANGE_TYPE, symbol, 5)
	cmdStatus := n.redisClient.Get(n.ctx, redisKey)
	if cmdStatus.Err() != nil {
		return ent1.Candle{}, redisKey, cmdStatus.Err()
	}

	bytes, _ := cmdStatus.Bytes()
	data, _, err := typeconv.UnmarshalCandleMsg(bytes)
	if err != nil {
		return ent1.Candle{}, redisKey, errors.Wrap(err, fmt.Sprintf("%s: %s", redisKey, "whalehunterRedisRepo.GetCandleData.redisClient.GetCandleData.msgp.Unmarshal"))
	}

	return data, redisKey, nil
}

func (n *redisRepo) GetCandlesData(symbol string, interval int) ([]ent1.Candle, int, string, error) {
	redisKey := fmt.Sprintf("%s:%s:%d:Candles", ent1.EXCHANGE_TYPE, symbol, interval)
	cmdStatus := n.redisClient.Get(n.ctx, redisKey)
	if cmdStatus.Err() != nil {
		return []ent1.Candle{}, 0, redisKey, cmdStatus.Err()
	}

	bytes, _ := cmdStatus.Bytes()
	data, maxCount, _, err := typeconv.UnmarshalCandlesMsg(bytes)
	if err != nil {
		return []ent1.Candle{}, 0, redisKey, errors.Wrap(err, fmt.Sprintf("%s: %s", redisKey, "whalehunterRedisRepo.GetCandlesData.redisClient.GetCandlesData.msqp.Unmarshal"))
	}
	return data, maxCount, redisKey, nil
}

func (n *redisRepo) GetDepthData(symbol string) (ent1.DepthData, string, error) {
	redisKey := fmt.Sprintf("%s:%s:Depth", ent1.EXCHANGE_TYPE, symbol)
	cmdStatus := n.redisClient.Get(n.ctx, redisKey)
	if cmdStatus.Err() != nil {
		return ent1.DepthData{}, redisKey, cmdStatus.Err()
	}

	bytes, _ := cmdStatus.Bytes()
	data, _, err := typeconv.UnmarshalDepthDataMsg(bytes)
	if err != nil {
		return ent1.DepthData{}, redisKey, errors.Wrap(err, fmt.Sprintf("%s: %s", redisKey, "whalehunterRedisRepo.GetDepthData.redisClient.GetDepthData.msqp.Unmarshal"))
	}

	return data, redisKey, nil
}

// SetCache Setting cache
func (n *redisRepo) SetCache(key string, value any) error {
	if err := n.redisClient.Set(n.ctx, key, value, 0).Err(); err != nil {
		return errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.SetCache")
	}

	return nil
}

// SetCache Setting cache
func (n *redisRepo) GetOpenCandlesCache(key string) (candles []ent1.Candle, err error) {
	var jsonRes []byte
	cmd := n.redisClient.Get(n.ctx, key)
	if err = cmd.Err(); err != nil {
		return nil, errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.GetCache")
	}
	jsonRes, err = cmd.Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.Bytes")
	}
	err = json.Unmarshal(jsonRes, &candles)
	if err != nil {
		return nil, errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.JSON.Unmarshal")
	}
	return
}

// SetCache Setting cache
func (n *redisRepo) GetOpenTickersCache(key string) (candles []ent1.TickCandle, err error) {
	var jsonRes []byte
	cmd := n.redisClient.Get(n.ctx, key)
	if err = cmd.Err(); err != nil {
		return nil, errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.GetCache")
	}
	jsonRes, err = cmd.Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.Bytes")
	}
	err = json.Unmarshal(jsonRes, &candles)
	if err != nil {
		return nil, errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.JSON.Unmarshal")
	}
	return
}

func (n *redisRepo) GetTickerData(symbol string) (ent1.TickCandle, string, error) {
	var candles []ent1.TickCandle
	var data ent1.TickCandle
	var err error
	var jsonRes []byte

	redisKey := fmt.Sprintf("%s:%s:10:ticker", ent1.EXCHANGE_TYPE, symbol)
	cmd := n.redisClient.Get(n.ctx, redisKey)
	if err = cmd.Err(); err != nil {
		return ent1.TickCandle{}, redisKey, errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.Get")
	}
	jsonRes, err = cmd.Bytes()
	if err != nil {
		return ent1.TickCandle{}, redisKey, errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.Bytes")
	}
	err = json.Unmarshal(jsonRes, &candles)
	if err != nil {
		return ent1.TickCandle{}, redisKey, errors.Wrap(err, "privatedataRedisRepo.SetCache.redisClient.JSON.Unmarshal")
	}
	data = candles[len(candles)-1]
	return data, redisKey, nil
}

// DeleteCache Delete cache
func (n *redisRepo) DeleteCache(key string) error {
	if err := n.redisClient.Del(n.ctx, key).Err(); err != nil {
		return errors.Wrap(err, "privatedataRedisRepo.DeleteCache.redisClient.Del")
	}

	return nil
}
