import unittest
import api
import models


class TestApi(unittest.TestCase):

    def test_get_exchange_info(self):
        api.get_exchange_info()

    def test_get_ticker(self):
        api.get_ticker()

    def test_get_depth(self):
        symbol = "xxx"
        size = "xxx"
        api.get_depth(symbol, size)

    def test_get_kline(self):
        symbol = "xxx"
        period = "xxx"
        size = "xxx"
        api.get_kline(symbol, period, size)

    def test_get_user_asset(self):
        api.get_user_asset()

    def test_place_limit_order(self):
        params = models.PlaceOrderRequestParams
        params.market = "xxx"
        params.side = "xxx"
        params.amount = "xxx"
        params.price = "xxx"

        api.place_limit_order(params)

    def test_place_market_order(self):
        params = models.PlaceOrderRequestParams
        params.market = "xxx"
        params.side = "xxx"
        params.amount = "xxx"

        api.place_market_order(params)

    def test_get_pending_order(self):
        params = models.GetPendingRequestParams
        params.market = "xxx"
        params.offset = "xxx"
        params.limit = "xxx"

        api.get_pending_order(params)

    def test_cancel_order(self):
        params = models.CancelRequestParams
        params.market = "xxx"
        params.order_id = "xxx"

        api.cancel_order(params)

    def test_get_order_deals(self):
        params = models.GetOrderDealsParams
        params.order_id = "xxx"
        params.limit = "xxx"
        params.offset = "xxx"

        api.get_order_deals(params)

    def test_get_finished_order(self):
        params = models.GetFinishedOrderParams
        params.market = "xxx"
        params.offset = "xxx"
        params.limit = "xxx"
        params.start_time = "xxx"
        params.end_time = "xxx"
        params.side = "xxx"

        api.get_finished_order(params)