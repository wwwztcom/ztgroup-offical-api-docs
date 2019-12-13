package exchange_api_demo_golang

import "testing"

// 获取交易对信息
func TestGetExchangeInfo(t *testing.T) {
	GetExchangeInfo()
}

// 获取Ticker
func TestGetTicker(t *testing.T) {

	GetTicker()
}

// 获取交易深度信息
func TestGetDepth(t *testing.T) {
	strSymbol := "ETH_CNT"
	size := "10"

	GetDepth(strSymbol, size)
}

// 获取Kline 数据
func TestGetKLine(t *testing.T) {
	strSymbol := "ETH_CNT"
	ty := "5min"
	size := "10"

	GetKLine(strSymbol, ty, size)
}

// 获取用户资产
func TestGetUserAsset(t *testing.T) {

	GetUserAsset()

}

// 下限价单
func TestPlaceLimitOrder(t *testing.T) {

	params := new(PlaceRequestParams)
	params.Price = "10000"
	params.Market = "ETH_CNT"
	params.Side = "2"
	params.Amount = "1"

	PlaceLimitOrder(params)
}

// 下市价单
func TestPlaceMarketOrder(t *testing.T) {

	params := new(PlaceRequestParams)

	params.Market = "ETH_CNT"
	params.Side = "2"
	params.Amount = "20"

	PlaceMarketOrder(params)
}

// 获取未成交订单
func TestGetPendingOrder(t *testing.T) {

	params := new(GetPendingRequestParams)

	params.Market = "ETH_CNT"
	params.Offset = "0"
	params.Limit = "5"

	GetPendingOrder(params)
}

// 取消订单
func TestCancelOrder(t *testing.T) {

	params := new(CancelRequestParams)

	params.Market = "ETH_CNT"
	params.Order_id = "4866125"

	CancelOrder(params)
}

// 获取已成交订单
func TestGetFinishedOrder(t *testing.T) {

	params := new(GetFinishedOrderParams)

	params.Market = "ETH_CNT"
	params.Start_time = "0"
	params.End_time = "0"
	params.Offset = "0"
	params.Limit = "5"
	params.Side = "0"

	GetFinishedOrder(params)
}
