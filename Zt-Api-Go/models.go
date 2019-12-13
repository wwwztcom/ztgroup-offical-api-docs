package exchange_api_demo_golang

// 获取用户资产参数
type AssetRequestParams struct {
}

// 下单参数
type PlaceRequestParams struct {
	Amount string                 // 限价表示下单数量, 市价买单时表示买多少钱, 市价卖单时表示卖多少币
	Price  string                 // 下单价格, 市价单不传该参数
	Market string                 // 交易对,
	Side   string                 // 方向, 1 卖 , 2 买
}

// 查询未成交订单参数
type GetPendingRequestParams struct {
	Market string                // 交易对,
	Offset string				 // 偏移
	Limit  string                // 限制 [1,100]
}

// 取消订单参数
type CancelRequestParams struct {
	Market   string             // 交易对
	Order_id string             // 订单ID
}

// 查询已成交订单参数
type GetFinishedOrderParams struct {
	Market     string 			// 交易对
	Start_time string           // 开始时间，以秒计数的时间戳，不限为0
	End_time   string			// 结束时间，以秒计数的时间戳，不限为0
	Offset     string			// 偏移
	Limit      string			// 限制 [1,100]
	Side       string			// 方向, 1 卖 , 2 买, 0 不限
}
