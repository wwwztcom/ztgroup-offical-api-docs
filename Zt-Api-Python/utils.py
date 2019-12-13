#!/usr/bin/python
# -*- coding: utf-8 -*-
# 用于进行http请求，以及MD5加密，生成签名的工具类

import requests
import models
import hashlib

# 构造request header
header = {
    "X-SITE-ID": models.X_SITE_ID,
    "Content-Type": models.ContentType,
}


# Http Get请求基础函数, 通过封装python语言Http请求, 支持ZT网REST API的HTTP Get请求
# url: 请求的URL
# return: 请求结果
def http_get_request(url: str):
    resp = requests.get(url, headers=header)
    return resp.text


# Http POST请求基础函数, 通过封装python语言Http请求, 支持ZT网REST API的HTTP POST请求
# strUrl: 请求的URL
# return: 请求结果
def http_post_request(url: str, params: dict):
    params["api_key"] = models.API_KEY
    params["sign"] = get_sign_data(params, models.SECRET_KEY)

    resp = requests.post(url, params, headers=header)

    return resp.text


# 获取签名参数
def get_sign_data(params: dict, secret_key: str):

    key_list = []
    sign_str = ""

    for k in params.keys():
        key_list.append(k)

    key_list.sort()

    for _, v in enumerate(key_list):
        sign_str = sign_str + v + "=" + params[v] + "&"

    sign_str = sign_str + "secret_key=" + secret_key
    sign = hashlib.md5(sign_str.encode("utf-8")).hexdigest().upper()

    return sign

