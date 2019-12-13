#!/usr/bin/python
# -*- coding: utf-8 -*-


# key
API_KEY = "xxx"
SECRET_KEY = "xxx"
# content-y=type
ContentType = "application/x-www-form-urlencoded"
# url
Host_Url = "https://www.zt.com"
# 站点ID
X_SITE_ID = "1"


# 下单参数
class PlaceOrderRequestParams(object):
    def __init__(self, amount, price, market, side):
        self.amount = amount
        self.price = price
        self.market = market
        self.side = side


# 查询未成交订单参数
class GetPendingRequestParams(object):
    def __init__(self, market, offset, limit):
        self.market = market
        self.offset = offset
        self.limit = limit


# 取消订单参数
class CancelRequestParams(object):
    def __init__(self, market, order_id):
        self.market = market
        self.order_id = order_id


# 查询订单成交详情参数
class GetOrderDealsParams(object):
    def __init__(self, order_id, offset, limit):
        self.order_id = order_id
        self.offset = offset
        self.limit = limit


# 查询已成交订单参数
class GetFinishedOrderParams(object):
    def __init__(self, market, offset, limit, start_time, end_time, side):
        self.market = market
        self.offset = offset
        self.limit = limit
        self.start_time = start_time
        self.end_time = end_time
        self.side = side
