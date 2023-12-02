package binance

import "github.com/bullean-ai/hexa-neural-net/pkg/utils/typeconv"

func ChartDataParser(arr [][]interface{}) (Linedata []ChartData) {
	for i := 0; i < len(arr); i++ {
		data := ChartData{
			OpenTime:   typeconv.ToInt64(arr[i][0]),
			OpenPrice:  typeconv.ToFloat(arr[i][1]),
			HighPrice:  typeconv.ToFloat(arr[i][2]),
			LowPrice:   typeconv.ToFloat(arr[i][3]),
			ClosePrice: typeconv.ToFloat(arr[i][4]),
			Volume:     typeconv.ToFloat(arr[i][5]),
			CloseTime:  typeconv.ToInt64(arr[i][6]),
		}
		Linedata = append(Linedata, data)
	}
	return
}
