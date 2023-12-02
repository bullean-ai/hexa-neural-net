package binance

import (
	"net/http"
)

const (
	PROVIDER_NAME = "binance"
)

var (
	httpResponse *http.Response
	Record       int64
	err          error
	Result       string
	arr          [][]interface{}
)

/*
func GetKlineData(cfg *config.Config, pair, interval string, limit int) (Linedata []ChartData) {
	URI := fmt.Sprintf("%s/%s/klines?symbol=%s&interval=%s&limit=%d", cfg.Binance.API_URL, cfg.Binance.API_VER, pair, interval, limit)
	fmt.Println(URI)
	httpResponse, err = http.Get(URI)
	if err != nil {
		fmt.Println(err)
		Record = -1
	}
	defer httpResponse.Body.Close()

	body, err := io.ReadAll(httpResponse.Body)
	err = json.Unmarshal(body, &arr)
	if err != nil {
		println(err.Error())
	}
	Linedata = ChartDataParser(arr)

	return
}

func GetFutureKlineData(cfg *config.Config, pair, interval string, limit int) (Linedata []ChartData) {
	//keyName := fmt.Sprintf("%s-%s_%s", PROVIDER_NAME, dat.Pair, dat.Interval)

	URI := fmt.Sprintf("%s/%s/klines?symbol=%s&interval=%s&limit=%d", cfg.Binance.FAPI_URL, cfg.Binance.FAPI_VER, pair, interval, limit)
	fmt.Printf("URIS : %s \n", URI)
	httpResponse, err = http.Get(URI)
	if err != nil {
		fmt.Println(err)
		Record = -1
	}
	defer httpResponse.Body.Close()

	body, err := io.ReadAll(httpResponse.Body)
	err = json.Unmarshal(body, &arr)
	if err != nil {
		println(err.Error())
	}
	Linedata = ChartDataParser(arr)
	return
}

func GetTakerLongShortRatioData(cfg *config.Config, symbol, period string) (Linedata []TakerLongShortRatioData) {

	URI := fmt.Sprintf("%s?symbol=%s&period=%s", cfg.Binance.FTAKER_LONG_SHORT_RATIO_URL, symbol, period)
	fmt.Printf("URIS : %s \n", URI)
	httpResponse, err = http.Get(URI)
	if err != nil {
		fmt.Println(err)
		Record = -1
	}
	defer httpResponse.Body.Close()

	body, err := io.ReadAll(httpResponse.Body)
	err = json.Unmarshal(body, &Linedata)
	if err != nil {
		println(err.Error())
	}

	return
}

func GetDepthData(cfg *config.Config, pair string, limit int) (LineData DepthData) {

	URI := fmt.Sprintf("%s/%s/depth?symbol=%s&limit=%d", cfg.Binance.API_URL, cfg.Binance.FAPI_VER, pair, limit)
	fmt.Printf("URIS : %s \n", URI)
	httpResponse, err = http.Get(URI)
	if err != nil {
		fmt.Println(err)
		Record = -1
	}
	defer httpResponse.Body.Close()

	body, err := io.ReadAll(httpResponse.Body)
	err = json.Unmarshal(body, &LineData)
	if err != nil {
		println(err.Error())
	}
	return
}

func GetFutureDepthData(cfg *config.Config, pair string, limit int) (LineData DepthData) {

	URI := fmt.Sprintf("%s/%s/depth?symbol=%s&limit=%d", cfg.Binance.FAPI_URL, cfg.Binance.FAPI_VER, pair, limit)
	fmt.Printf("URIS : %s \n", URI)
	httpResponse, err = http.Get(URI)
	if err != nil {
		fmt.Println(err)
		Record = -1
	}
	defer httpResponse.Body.Close()

	body, err := io.ReadAll(httpResponse.Body)
	err = json.Unmarshal(body, &LineData)
	if err != nil {
		println(err.Error())
	}
	return
}

func GetPriceTickerData(cfg *config.Config, pairs []string) (arrPTicker []PriceTickerData) {

	pairsByte, err := json.Marshal(pairs)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	params := url.Values{}
	jsonStr := string(pairsByte)
	params.Add("symbols", jsonStr)

	URI := fmt.Sprintf("%s/%s/ticker/24hr?%s", cfg.Binance.API_URL, cfg.Binance.API_VER, params.Encode())
	fmt.Printf("URIS : %s \n", URI)
	httpResponse, err = http.Get(URI)
	if err != nil {
		fmt.Println(err)
		Record = -1
	}
	defer httpResponse.Body.Close()

	body, err := io.ReadAll(httpResponse.Body)
	err = json.Unmarshal(body, &arrPTicker)
	if err != nil {
		println(err.Error())
	}

	return
}
*/
