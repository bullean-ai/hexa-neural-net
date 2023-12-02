package typeconv

import (
	"fmt"
	entities "github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/entities"
	"github.com/tinylib/msgp/msgp"
	"strconv"
	"time"
)

func ToInt(key interface{}) int {
	v, err := strconv.Atoi((key).(string))
	if err != nil {
		return 0
	}
	return v
}

func ToFloat(key interface{}) float64 {
	v, err := strconv.ParseFloat(key.(string), 64)
	if err != nil {
		return 0
	}
	return v
}

func ToInt64(key interface{}) int64 {
	v := int64(key.(float64))
	return v
}

func IChkStr(obj interface{}) string {
	if str, ok := obj.(string); ok {
		return str
	} else {
		return ""
	}
}

func IChkF64(obj interface{}) float64 {
	if str, ok := obj.(float64); ok {
		return str
	} else {
		return 0.0
	}
}

func IChkF64s(obj interface{}) string {
	if str, ok := obj.(float64); ok {
		//return strconv.FormatFloat(str, '', -1, 64)
		return fmt.Sprintf("%.8f", str)
	} else {
		return ""
	}
}

func IChkI64(obj interface{}) int64 {
	if str, ok := obj.(int64); ok {
		return str
	} else {
		return 0
	}
}

// UnmarshalCandleMsg implementitiess msgp.Unmarshaler
func UnmarshalCandleMsg(bytes []byte) (tick entities.Candle, data []byte, err error) {
	var zb0001 uint32
	zb0001, bytes, err = msgp.ReadArrayHeaderBytes(bytes)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 9 {
		err = msgp.ArrayError{Wanted: 9, Got: zb0001}
		return
	}
	if msgp.IsNil(bytes) {
		bytes, err = msgp.ReadNilBytes(bytes)
		if err != nil {
			return
		}
		tick.Date = nil
	} else {
		if tick.Date == nil {
			tick.Date = new(time.Time)
		}
		*tick.Date, bytes, err = msgp.ReadTimeBytes(bytes)
		if err != nil {
			err = msgp.WrapError(err, "Date")
			return
		}
	}
	tick.Open, bytes, err = msgp.ReadFloat64Bytes(bytes)
	if err != nil {
		err = msgp.WrapError(err, "Open")
		return
	}
	tick.High, bytes, err = msgp.ReadFloat64Bytes(bytes)
	if err != nil {
		err = msgp.WrapError(err, "High")
		return
	}
	tick.Low, bytes, err = msgp.ReadFloat64Bytes(bytes)
	if err != nil {
		err = msgp.WrapError(err, "Low")
		return
	}
	tick.Close, bytes, err = msgp.ReadFloat64Bytes(bytes)
	if err != nil {
		err = msgp.WrapError(err, "Close")
		return
	}
	tick.Volume, bytes, err = msgp.ReadFloat64Bytes(bytes)
	if err != nil {
		err = msgp.WrapError(err, "Volume")
		return
	}
	tick.QuoteAssetVolume, bytes, err = msgp.ReadFloat64Bytes(bytes)
	if err != nil {
		err = msgp.WrapError(err, "QuoteAssetVolume")
		return
	}
	tick.TakerBaseVolume, bytes, err = msgp.ReadFloat64Bytes(bytes)
	if err != nil {
		err = msgp.WrapError(err, "TakerBaseVolume")
		return
	}
	tick.TakerQuoteVolume, bytes, err = msgp.ReadFloat64Bytes(bytes)
	if err != nil {
		err = msgp.WrapError(err, "TakerQuoteVolume")
		return
	}
	data = bytes
	return
}

// UnmarshalCandlesMsg implementitiess msgp.Unmarshaler
func UnmarshalCandlesMsg(data []byte) (candles []entities.Candle, maxCount int, bytes []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, data, err = msgp.ReadMapHeaderBytes(data)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, data, err = msgp.ReadMapKeyZC(data)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Candles":
			var zb0002 uint32
			zb0002, data, err = msgp.ReadArrayHeaderBytes(data)
			if err != nil {
				err = msgp.WrapError(err, "Candles")
				return
			}
			if cap(candles) >= int(zb0002) {
				candles = (candles)[:zb0002]
			} else {
				candles = make([]entities.Candle, zb0002)
			}
			for i := range candles {
				candles[i], data, err = UnmarshalCandleMsg(data)
				if err != nil {
					err = msgp.WrapError(err, "Candles", i)
					return
				}
			}
		case "MaxCount":
			maxCount, data, err = msgp.ReadIntBytes(data)
			if err != nil {
				err = msgp.WrapError(err, "MaxCount")
				return
			}
		default:
			data, err = msgp.Skip(data)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// UnmarshalDepthDataMsg implementitiess msgp.Unmarshaler
func UnmarshalDepthDataMsg(bts []byte) (depth entities.DepthData, o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Symbol":
			depth.Symbol, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Symbol")
				return
			}
		case "Bids":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Bids")
				return
			}
			if cap(depth.Bids) >= int(zb0002) {
				depth.Bids = (depth.Bids)[:zb0002]
			} else {
				depth.Bids = make([]entities.DepthPrice, zb0002)
			}
			for za0001 := range depth.Bids {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
				} else {
					depth.Bids[za0001], bts, err = UnmarshalDepthPriceMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Bids", za0001)
						return
					}
				}
			}
		case "Asks":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Asks")
				return
			}
			if cap(depth.Asks) >= int(zb0003) {
				depth.Asks = (depth.Asks)[:zb0003]
			} else {
				depth.Asks = make([]entities.DepthPrice, zb0003)
			}
			for za0002 := range depth.Asks {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
				} else {
					depth.Asks[za0002], bts, err = UnmarshalDepthPriceMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Asks", za0002)
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// UnmarshalDepthPriceMsg implements msgp.Unmarshaler
func UnmarshalDepthPriceMsg(bts []byte) (depthPrice entities.DepthPrice, o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Price":
			depthPrice.Price, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Price")
				return
			}
		case "Sale":
			depthPrice.Sale, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Sale")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}
