import unittest
import ws


class TestApi(unittest.TestCase):
    def test_sub_client(self):
        ws.sub_client()

    def test_query_client(self):
        ws.query_client()
