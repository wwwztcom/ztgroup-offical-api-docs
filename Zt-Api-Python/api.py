import utils
import models


# 查询系统支持的所有币种,交易精度
def get_exchange_info():
    req_path = "/api/v1/exchangeInfo"
    str_url = models.Host_Url + req_path
    exchange_info = utils.http_get_request(str_url)

    print(exchange_info)


# 获取聚合行情
def get_ticker():
    req_path = "/api/v1/tickers"
    str_url = models.Host_Url + req_path
    ticker = utils.http_post_request(str_url)

    print(ticker)


# 获取交易深度信息
# symbol: 交易对, BTC_USDT
# size: 获取数量
def get_depth(symbol: str, size: str):
    req_path = "/api/v1/depth?symbol={}&size={}".format(symbol, size)
    str_url = models.Host_Url + req_path
    depth = utils.http_get_request(str_url)

    print(depth)


# 获取K线数据
# symbol: 交易对, BTC_USDT
# period: K线类型, 1min, 5min, 15min......
# size: 获取数量, [1-2000]
def get_kline(symbol: str, period: str, size: str):
    req_path = "/api/v1/kline?symbol={}&type={}&size={}".format(symbol, period, size)
    str_url = models.Host_Url + req_path
    kline = utils.http_get_request(str_url)

    print(kline)


# 获取用户资产
def get_user_asset():
    params_dist = {}
    str_url = models.Host_Url + "/api/v1/private/user"
    user_asset = utils.http_post_request(str_url, params_dist)

    print(user_asset)


# 挂限价单
def place_limit_order(params):
    params_dist = {
        "price": params.price,
        "market": params.market,
        "side": params.side,
        "amount": params.amount,
    }

    str_url = models.Host_Url + "/api/v1/private/trade/limit"
    place_order_return = utils.http_post_request(str_url, params_dist)

    print(place_order_return)


# 挂市价单
def place_market_order(params):
    params_dist = {
        "market": params.market,
        "side": params.side,
        "amount": params.amount,
    }

    str_url = models.Host_Url + "/api/v1/private/trade/market"
    place_order_return = utils.http_post_request(str_url, params_dist)

    print(place_order_return)


# 获取未成交订单
def get_pending_order(params):
    params_dist = {
        "market": params.market,
        "offset": params.offset,
        "limit": params.limit,
    }

    str_url = models.Host_Url + "/api/v1/private/order/pending"
    pending_order_return = utils.http_post_request(str_url, params_dist)

    print(pending_order_return)


# 取消订单
def cancel_order(params):
    params_dist = {
        "order_id": params.order_id,
        "market": params.market,
    }

    str_url = models.Host_Url + "/api/v1/private/trade/cancel"
    cancel_order_return = utils.http_post_request(str_url, params_dist)

    print(cancel_order_return)


# 􏱐􏱑􏱨􏱩􏰼􏰠􏱐􏱑􏱨􏱩􏰼􏰠􏱐􏱑􏱨􏱩􏰼􏰠􏱐􏱑􏱨􏱩􏰼􏰠查询订单成交详情
def get_order_deals(params):
    params_dist = {
        "order_id": params.order_id,
        "offset": params.offset,
        "limit": params.limit,
    }

    str_url = models.Host_Url + "/api/v1/private/order/deals"
    order_deals_return = utils.http_post_request(str_url, params_dist)

    print(order_deals_return)


# 获取已成交订单
def get_finished_order(params):
    params_dist = {
        "market": params.market,
        "start_time": params.start_time,
        "end_time": params.end_time,
        "offset": params.offset,
        "limit": params.limit,
        "side": params.side,
    }

    str_url = models.Host_Url + "/api/v1/private/order/finished"
    finished_order_return = utils.http_post_request(str_url, params_dist)

    print(finished_order_return)
