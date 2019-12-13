# !/usr/bin/env python3
# coding:utf-8

"""
    @author: Lmj
    @date: 2019/12/13
"""

from websocket import create_connection
from threading import Thread
import json
import time
import random


# example
# 市场状态订阅
def sub_client():
    url = "wss://ws.zt.com/ws"
    conn = create_connection(url)
    rand_num = random.randint(1, 10000)
    msg = json.dumps({"method": "state.subscribe", "params": ["BTC_USDT"], "id": rand_num})
    conn.send(msg)

    # 启用新线程
    t = Thread(target=ping, args=(conn,))
    t.start()

    while True:
        rec_msg = conn.recv()
        print(rec_msg)


def query_client():
    url = "wss://ws.zt.com/ws"
    conn = create_connection(url)
    rand_num = random.randint(1, 10000)
    msg = json.dumps({"method": "state.query", "params": ["BTC_USDT", 86400], "id": rand_num})
    conn.send(msg)

    rec_msg = conn.recv()
    print(rec_msg)

    conn.close()


# 客户端向服务器发送心跳信息
def ping(conn):
    msg = json.dumps({"method": "server.ping", "params": [], "id": 10086})

    while True:
        conn.send(msg)
        time.sleep(15)
