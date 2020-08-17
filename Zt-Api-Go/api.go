package exchange_api_demo_golang

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

const (
	API_KEY     = "" // "xxx"
	SECRET_KEY  = "" // "xxx"
	ContentType = "application/x-www-form-urlencoded"
	HostUrl     = "https://www.ztb.com" // "xxx"
	X_SITE_ID   = "1" // "xxx"
)

// 查询系统支持的所有币种,交易精度
func GetExchangeInfo() {

	strRequestUrl := "/api/v1/exchangeInfo"
	strUrl := HostUrl + strRequestUrl

	exchangeInfoReturn, err := HttpGetRequest(strUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(exchangeInfoReturn))
}

// 获取聚合行情
func GetTicker() {

	strRequestUrl := "/api/v1/tickers"
	strUrl := HostUrl + strRequestUrl
	logs.Info(strUrl)
	tickerReturn, err := HttpGetRequest(strUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(tickerReturn))
}

// 获取交易深度信息
// strSymbol: 交易对, BTC_USDT
// size: 获取数量
func GetDepth(strSymbol string, size string) {

	strRequestUrl := fmt.Sprintf("/api/v1/depth?symbol=%s&size=%s", strSymbol, size)
	strUrl := HostUrl + strRequestUrl

	depthReturn, err := HttpGetRequest(strUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(depthReturn))
}

// 获取K线数据
// strSymbol: 交易对, BTC_USDT
// strPeriod: K线类型, 1min, 5min, 15min......
// size: 获取数量, [1-2000]
func GetKLine(strSymbol, strPeriod, size string) {

	strRequestUrl := fmt.Sprintf("/api/v1/kline?symbol=%s&type=%s&size=%s", strSymbol, strPeriod, size)
	strUrl := HostUrl + strRequestUrl

	klineReturn, err := HttpGetRequest(strUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(klineReturn)
}

// 获取用户资产
func GetUserAsset() {

	mapParams := make(map[string]string)

	strRequestUrl := HostUrl + "/api/v1/private/user"
	userAssetReturn, err := HttpPostRequest(mapParams, strRequestUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(userAssetReturn)
}

// 挂限价单
func PlaceLimitOrder(params *PlaceRequestParams) {

	mapParams := make(map[string]string)
	mapParams["market"] = params.Market
	mapParams["side"] = params.Side
	mapParams["amount"] = params.Amount
	mapParams["price"] = params.Price

	strRequestUrl := HostUrl + "/api/v1/private/trade/limit"
	placeOrderReturn, err := HttpPostRequest(mapParams, strRequestUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(placeOrderReturn)
}

// 下市价单
func PlaceMarketOrder(params *PlaceRequestParams) {

	mapParams := make(map[string]string)
	mapParams["market"] = params.Market
	mapParams["side"] = params.Side
	mapParams["amount"] = params.Amount

	strRequestUrl := HostUrl + "/api/v1/private/trade/market"
	placeOrderReturn, err := HttpPostRequest(mapParams, strRequestUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(placeOrderReturn)
}

// 获取未成交订单
func GetPendingOrder(params *GetPendingRequestParams)  {

	mapParams := make(map[string]string)
	mapParams["market"] = params.Market
	mapParams["offset"] = params.Offset
	mapParams["limit"] = params.Limit

	strRequestUrl := HostUrl + "/api/v1/private/order/pending"
	pendingOrderReturn, err := HttpPostRequest(mapParams, strRequestUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pendingOrderReturn)
}

// 取消订单
func CancelOrder(params *CancelRequestParams) {

	mapParams := make(map[string]string)
	mapParams["market"] = params.Market
	mapParams["order_id"] = params.Order_id

	strRequestUrl := HostUrl + "/api/v1/private/trade/cancel"
	cancelReturn, err := HttpPostRequest(mapParams, strRequestUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cancelReturn)
}


// 获取已成交订单
func GetFinishedOrder(params *GetFinishedOrderParams) {

	mapParams := make(map[string]string)
	mapParams["market"] = params.Market
	mapParams["start_time"] = params.Start_time
	mapParams["end_time"] = params.End_time
	mapParams["offset"] = params.Offset
	mapParams["limit"] = params.Limit
	mapParams["side"] = params.Side

	strRequestUrl := HostUrl + "/api/v1/private/order/finished"
	finishedReturn, err := HttpPostRequest(mapParams, strRequestUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(finishedReturn)
}



