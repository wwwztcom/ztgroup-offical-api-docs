# ZT OPEN API

简体中文| [English](./readme-en.md)

- [**入门指引**](#入门指引)
  
  - [**创建API Key**](#创建api-key)
  - [**接口调用方式说明**](#接口调用方式说明)
- [**服务器**](#服务器)
  
- [**REST API**](#rest-api)
  - [**接入 URL**](#接入-url)
  - [**请求交互**](#请求交互)
  - [**！！注意！！**](#注意)
  - [**获取交易所市场数据**](#获取交易所市场数据)
  - [**获取深度信息**](#获取深度信息)
  - [**获取最新成交记录**](#获取最新成交记录)
  - [**获取K线信息**](#获取K线信息)
  - [**获取交易对信息**](#获取交易对信息)
  - [**签名认证**](#签名认证)
  - [**获取用户资产**](#获取用户资产)
  - [**限价交易**](#限价交易)
  - [**市价交易**](#市价交易)
  - [**取消某个委托订单**](#取消某个委托订单)
  - [**批量取消委托订单**](#批量取消委托订单)
  - [**查询未成交订单**](#查询未成交订单)
  - [**查询某个未成交订单详情**](#查询某个未成交订单详情)
  - [**查询已成交订单接口**](#查询已成交订单接口)
  - [**查询某个已成交订单详情**](#查询某个已成交订单详情)
  - [**查询订单成交接口**](#查询订单成交接口(分笔交明细))

- [**Websocket API**](#websocket-api)

  - [**简介**](#简介) 
  - [**市场状态(市场概要)**](#市场状态（市场概要）)
  - [**今日市场状态**](#今日市场状态（市场概要）)
  - [**K线数据**](#K线数据)
  - [**市场深度行情数据**](#市场深度行情数据)
  - [**市场最新价格数据**](#市场最新价格数据)
  - [**最新成交数据**](#最新成交数据)

  

## 入门指引

**欢迎使用开发者文档，ZT提供了简单易用的API接口，通过API可以获取市场行情数据、进行交易、管理订单**

### 创建API Key

用户在 **[ZT](https://www.ztb.im)** 注册账号后，需要在[API管理]中创建API Key秘钥，创建完成后得到一组随机生成的API Key与Secret Key,利用这一组数据可以进行程序化交易，单个账号最多创建5个密钥

> **请不要泄露API Key 与 Secret Key信息，以免造成资产损失,建议用户为API绑定IP地址，每个密钥最多绑定5个IP，使用英文逗号进行分隔**

### 接口调用方式说明

- REST API

  提供行情查询、余额查询、币币交易、订单管理功能，建议用户使用REST API进行账户余额查询、币币交易及订单管理等操作

### 服务器

- 服务器运行在东京，为了最大限度地减少API访问延迟，建议使用与东京通讯通畅的服务器


## REST API

### 接入 URL

- [https://www.ztb.im](https://www.ztb.im) 

### 请求交互

#### 介绍

REST API 提供行情查询、余额查询、币币交易、订单管理功能

所有请求基于Https协议，请求头信息中content-type需要统一设置为表单格式:

- **content-type:application/x-www-form-urlencoded**

#### 错误码

| 错误码   | 说明         | 原因             |
| :---- | :--------- | :------------- |
| 0     | 成功         |                |
| 1     | 参数不合法      |                |
| 2     | 内部错误       |                |
| 3     | 服务不可用      |                |
| 4     | 方法未找到      |                |
| 5     | 服务超时       |                |
| 10    | 金额不足       |                |
| 11    | 交易数量太小     |                |
| 12    | 深度不足       |                |
| 10005 | 记录未找到      | X-SITE-ID，设置错误 |
| 10022 | 用户未实名      |                |
| 10051 | 用户禁止交易     |                |
| 10056 | 小于最低金额     |                |
| 10059 | 该资产暂未开启交易  |                |
| 10060 | 该交易对暂未开启交易 |                |
| 10062 | 金额精度不正确    |                |

### 注意

- 所有接口请求（公共和私有接口）都必须在Request请求的Header中添加 X-SITE-ID 字段，该字段的值为”1“。该字段不用做签名校验。
- 调用频率限制根据UID限制，各个接口的限制有所差异。


# 行情API 公共接口

## 获取交易所市场数据

Get  /api/v1/tickers  

频率限制：20次/s

### 示例

```
Request:
GET https://www.ztb.im/api/v1/tickers

Response:
{
  "ticker": [
    {
      "buy": "28.3437",
      "change": "2.95",    
      "high": "29.8265",
      "last": "28.8402",
      "low": "27.9001",
      "sell": "29.3233",
      "symbol": "ZEC_USDT",
      "vol": "18564.4392"
    },
    ...
  ],
  "timestamp": "1577680809298"
}
```

### 返回数据说明
```
timestamp: 服务端时间戳
buy: 最佳BID
change: 涨跌幅
high: 最高价
last: 最新价
low: 最低价
sell: 最佳ASK
vol: 成交量 (24小时交易量)
```
---

## 获取深度信息

Get /api/v1/depth 

频率限制：20次/s

### 请求参数

| 参数     | 描述      |
| ------ | ------- |
| symbol | 市场名称    |
| size   | 返回深度的条数 |

### 示例
```
Request:
GET https://www.ztb.im/api/v1/depth?symbol=BTC_USDT&size=1

Response:
{
    "asks":[
        ["0.387","43189.58824906"]
    ],
    "bids":[
        ["0.378","2078.91534391"]
    ]
}
```

### 返回数据说明
```
asks : ask深度[价格，数量]
bids : bid深度[价格，数量]
```
---

## 获取最新成交记录

Get /api/v1/trades  

频率限制：20次/s

### 请求参数
| 参数     | 描述      |
| ------ | ------- |
| symbol | 市场名称    |
| size   | 返回深度的条数 |

### 示例

```
Request:
GET https://www.ztb.im/api/v1/trades?symbol=BTC_USDT&size=1

Response:
[
    {
        "amount":"500",
        "price":"0.401",
        "side":"sell",
        "timestamp":"1535507624521"
    },
    {
        "amount":"442",
        "price":"0.401",
        "side":"sell",
        "timestamp":"1535507612055"
    }
]
```

### 返回数据说明
```
amount : 数量
price : 价格
side: 买入或者卖出
timestamp: 时间戳
```

---

## 获取K线信息

Get /api/v1/kline  

频率限制：20次/s

### 请求参数

| 参数     | 描述                                       |
| ------ | ---------------------------------------- |
| symbol | 市场名称                                     |
| type   | 分时参数，可以为1min,5min,15min,30min,hour,day,week |
| size   | 返回成交数的条数                                 |

### 示例

```
Request:
GET https://www.ztb.im/api/v1/kline?symbol=BTC_USDT&type=1min&size=10

Response:
[
    [
        1535508060000,
        "0.401",
        "0.401",
        "0.401",
        "0.401",
        "0"
    ],
    [
        1535508120000,
        "0.401",
        "0.401",
        "0.401",
        "0.401",
        "0"
    ],
    [
        1535508180000,
        "0.401",
        "0.401",
        "0.401",
        "0.401",
        "0"
    ]
    ...
]
```

### 返回数据说明

```
1532655300000: 时间戳  
2370.16: 开
2380: 高
2352: 低
2367.37: 收 
17259.83: 交易量
```

---

## 获取交易对信息

Get /api/v1/exchangeInfo

频率限制：20次/s

### 示例
```
Request:
GET https://www.ztb.im/api/v1/exchangeInfo

Response:
[
    {
        "baseAsset":"BTC",
        "baseAssetPrecision":8,
        "quoteAsset":"USDT",
        "quoteAssetPrecision":8,
        "status":"trading",
        "symbol":"BTC_USDT"
    },
    {
        "baseAsset":"ETH",
        "baseAssetPrecision":8,
        "quoteAsset":"USDT",
        "quoteAssetPrecision":8,
        "status":"trading",
        "symbol":"ETH_USDT"
    }
]
```

### 返回数据说明

```
baseAsset: 基础货币 
baseAssetPrecision: 基础货币精度
quoteAsset: 计价货币
quoteAssetPrecision: 计价货币精度
status: 交易状态
symbol: 交易对

```

---
# 交易API 私有接口

- 所有私有接口请求都使用POST方法，参数以form-data的形式提交（所有私有接口api_key和sign都必须传递）。

## 签名认证

- 所有的参数都必须进行签名验证，所有参数必须根据字母表按照参数名进行排序

### 示例

```
请求参数：

'market':'BTC_USDT',
'side':1,
'price':'50',
'amount':'0.02'

排序后的参数字符串:amount=0.02&api_key=apiKey&market=BTC_USDT&price=50&side=1

注意:生成MD5签名必须要secret_key,在以上生成的字符串基础上添加secret_key以生成最终的字符串。

最终签名字符串:amount=0.02&api_key=apiKey&market=BTC_USDT&price=50&side=1&secret_key=secret_key

MD5签名：
使用32bit的MD5加密字符串，生成的加密字符串必须大写
```

---
## 获取用户资产

POST /api/v1/private/user 

频率限制：5次/s

### 示例

```
Request:
POST https://www.ztb.im/api/v1/private/user

Response:
{
  "code": 0,
  "message": "操作成功",
  "result": {
    "USDT": {
      "available": "10718.74453852",
      "freeze": "0.10999996",
      "other_freeze": "0",
      "recharge_status": 0,
      "trade_status": 1,
      "withdraw_fee": "0.1",
      "withdraw_max": "1000",
      "withdraw_min": "0.001",
      "withdraw_status": 1
    },
    "ETH": {
      "available": "395.196",
      "freeze": "0",
      "other_freeze": "0",
      "recharge_status": 1,
      "trade_status": 1,
      "withdraw_fee": "0.01",
      "withdraw_max": "100000",
      "withdraw_min": "0.001",
      "withdraw_status": 1
    },
    "BTC": {
      "available": "46.20370336",
      "freeze": "0",
      "other_freeze": "0",
      "recharge_status": 0,
      "trade_status": 1,
      "withdraw_fee": "0.1",
      "withdraw_max": "100000",
      "withdraw_min": "1",
      "withdraw_status": 1
    }
  }
}
```

###  返回数据说明

```
available: 可用余额
freeze: 交易冻结金额
other_freeze: 其他冻结金额（包括C2C和提币冻结）
recharge_status: 充值状态0为不可充值，1为可充值
withdraw_fee: 提币手续费
withdraw_max: 最大提币金额
withdraw_min: 最小提币金额
withdraw_status: 提币状态0为不可提币，1为可提币
```

---
## 限价交易

POST /api/v1/private/trade/limit  

频率限制：500次/s

### 请求参数

| 参数   | 描述                   | 类型   | 值                     |
| ------ | ---------------------- | ------ | ---------------------- |
| market | 市场                   | string | 自定义（eg: BTC_USDT） |
| side   | 1为ASK卖出，2为BID买入 | string | 1,2                    |
| amount | 数量                   | string | 自定义                 |
| price  | 价格                   | string | 自定义                 |

### 示例
```
Request:
POST https://www.ztb.im/api/v1/private/trade/limit

Response:
{
  "code": 0,
  "message": "操作成功",
  "result": {
    "amount": "1",
    "ctime": 1535537926.246487,
    "deal_fee": "0",
    "deal_money": "0",
    "deal_stock": "0",
    "id": 32865,
    "left": "1",
    "maker_fee": "0.001",
    "market": "BTC_USDT",
    "mtime": 1535537926.246487,
    "price": "10",
    "side": 2,
    "source": "web,1",
    "taker_fee": "0.001",
    "type": 1,
    "user": 670865
  }
}
```

### 返回数据说明

```
amount: 数量
ctime: 创建时间
deal_fee: 成交手续费
deal_money: 成交金额
deal_stock: 成交资产
id: 编号
left: 剩余
maker_fee: maker手续费
market: 市场名
mtime: 发布到市场时间
price: 价格
side: 1为ASK卖出，2为BID买入
source:来源
taker_fee: taker手续费
type: 交易类型，1为限价，2为市价
user: 用户编号
```

---
## 市价交易

POST /api/v1/private/trade/market  用户市价交易

频率限制：500次/s

### 请求参数

| 参数   | 描述                   | 类型   | 值                     |
| ------ | ---------------------- | ------ | ---------------------- |
| market | 市场                   | string | 自定义（eg: BTC_USDT） |
| side   | 1为ASK卖出，2为BID买入 | string | 1,2                    |
| amount | 数量                   | string | 自定义                 |

### 示例

```
Request: 
POST https://www.ztb.im/api/v1/private/trade/market

Response:
{
  "code": 0,
  "message": "操作成功",
  "result": {
    "amount": "1",
    "ctime": 1535538409.189721,
    "deal_fee": "0.00019607843",
    "deal_money": "0.999999993",
    "deal_stock": "0.19607843",
    "id": 32868,
    "left": "7.0000000e-9",
    "maker_fee": "0",
    "market": "BTC_USDT",
    "mtime": 1535538409.189735,
    "price": "0",
    "side": 2,
    "source": "web,1",
    "taker_fee": "0.001",
    "type": 2,
    "user": 670865
  }
}
```

### 返回数据说明

```
amount: 数量
ctime: 创建时间
deal_fee: 成交手续费
deal_money: 成交金额
deal_stock: 成交资产
id: 编号
left: 剩余
maker_fee: maker手续费
market: 市场名
mtime: 发布到市场时间
price: 价格
side: 1为ASK卖出，2为BID买入
source:来源
taker_fee: taker手续费
type: 交易类型，1为限价，2为市价
user: 用户编号
```

---
## 取消某个委托订单

POST /api/v1/private/trade/cancel 

频率限制：100次/s

### 请求参数
| 参数     | 描述     | 类型   | 值                     |
| -------- | -------- | ------ | ---------------------- |
| market   | 市场名称 | string | 自定义（eg: BTC_USDT） |
| order_id | 订单编号 | string | 自定义（eg: 32865）    |

### 示例

```
Request: 
POST https://www.ztb.im/api/v1/private/trade/cancel

Response:
{
  "code": 0,
  "message": "操作成功",
  "result": {
    "amount": "1",
    "ctime": 1535538409.189721,
    "deal_fee": "0.00019607843",
    "deal_money": "0.999999993",
    "deal_stock": "0.19607843",
    "id": 32868,
    "left": "7.0000000e-9",
    "maker_fee": "0",
    "market": "BTC_USDT",
    "mtime": 1535538409.189735,
    "price": "0",
    "side": 2,
    "source": "web,1",
    "taker_fee": "0.001",
    "type": 2,
    "user": 670865
  }
}
```

### 返回数据说明

```
amount: 数量
ctime: 创建时间
deal_fee: 成交手续费
deal_money: 成交金额
deal_stock: 成交资产
id: 编号
left: 剩余
maker_fee: maker手续费
market: 市场名
mtime: 发布到市场时间
price: 价格
side: 1为ASK卖出，2为BID买入
source:来源
taker_fee: taker手续费
type: 交易类型，1为限价，2为市价
user: 用户编号
```

---
## 批量取消委托订单

POST /api/v1/private/trade/cancel_batch 每次批量取消委托订单数量不超过10个。

频率限制：10次/s

### 请求参数

| 参数        | 描述     | 取值 | 取值                                                         |
| ----------- | -------- | ---- | ------------------------------------------------------------ |
| orders_json | 订单编号 | json | [{"market":"BTC_USDT", "order_id":456647},{"market":"BTC_USDT", "order_id":456648}] |

### 示例

```
Request:
POST https://www.ztb.im/api/v1/private/trade/cancel_batch

Response:
{
  "code": 0,
  "message": "操作成功",
  "result": [
    {
      "market": "BTC_USDT",
      "order_id": 458815,
      "result": true
    },
    {
      "market": "BTC_USDT",
      "order_id": 458813,
      "result": true
    },
    {
      "market": "BTC_USDT",
      "order_id": 458812,
      "result": false
    }
  ]
}
```

### 返回数据说明

```
market: 市场
order_id: 订单编号
result: 取消结果(true 表示取消成功，false 表示取消失败)
```

------
## 查询未成交订单

POST /api/v1/private/order/pending  

频率限制：20次/s

### 请求参数

| 参数   | 描述   | 类型   | 值                     |
| ------ | ------ | ------ | ---------------------- |
| market | 市场   | string | 自定义（eg: BTC_USDT） |
| offset | 偏移   | string | 自定义                 |
| limit  | 限制值 | string | 自定义（不超过100）    |


### 示例
```
# Request 
POST https://www.ztb.im/api/v1/private/order/pending
# Response
{
  "code": 0,
  "message": "操作成功",
  "result": {
    "limit": 10,
    "offset": 0,
    "records": [
      {
        "amount": "1",
        "ctime": 1535544362.168106,
        "deal_fee": "0",
        "deal_money": "0",
        "deal_stock": "0",
        "id": 32871,
        "left": "1",
        "maker_fee": "0.001",
        "market": "BTC_USDT",
        "mtime": 1535544362.168106,
        "price": "5.1",
        "side": 2,
        "source": "web,1",
        "status": 1,
        "taker_fee": "0.001",
        "type": 1,
        "user": 670865
      }
    ],
    "total": 1
  }
}
```

### 返回数据说明

```
amount: 数量
ctime: 创建时间
deal_fee: 成交手续费
deal_money: 成交金额
deal_stock: 成交资产
id: 编号
left: 剩余
maker_fee: maker手续费
market: 市场名
mtime: 发布到市场时间
price: 价格
side: 1为ASK卖出，2为BID买入
source:来源
status: 1-初始化，2-已触发，3-已取消，4-部分成交，5-完全成交
taker_fee: taker手续费
type: 交易类型，1为限价，2为市价
user: 用户编号
```

------
## 查询某个未成交订单详情

POST /api/v1/private/order/pending/detail  

频率限制：10次/s

### 请求参数

| 参数     | 描述   | 类型   |                        |
| -------- | ------ | ------ | ---------------------- |
| market   | 市场   | string | 自定义（eg: BTC_USDT） |
| Order_id | 订单号 | string | 自定义（eg:  1080）    |

### 示例
```
Request:
POST https://www.ztb.im/api/v1/private/order/pending/detail

Response:
{
  "code": 0,
  "message": "操作成功",
  "result": {
    "amount": "10",
    "ctime": 1565681852.879657,
    "deal_fee": "0",
    "deal_money": "0",
    "deal_stock": "0",
    "id": 1080,
    "left": "10",
    "maker_fee": "0",
    "market": "BTC_USDT",
    "mtime": 1565681852.879657,
    "price": "1",
    "side": 2,
    "source": "web,127",
    "taker_fee": "0",
    "type": 1,
    "user": 2
  }
}
```

### 返回数据说明

```
amount: 数量
ctime: 创建时间
deal_fee: 成交手续费
deal_money: 成交金额
deal_stock: 成交资产
id: 编号
left: 剩余
maker_fee: maker手续费
market: 市场名
ftime: 发布到市场时间
price: 价格
side: 1为ASK卖出,2为BID买入
source: 来源
taker_fee: taker手续费
type: 交易类型,1为限价,2为市价
user: 用户编号
```

---
## 查询已成交订单接口

POST /api/v1/private/order/finished 

频率限制：10次/s

### 请求参数

| 参数       | 描述                                | 类型   | 值                     |
| ---------- | ----------------------------------- | ------ | ---------------------- |
| market     | 市场                                | string | 自定义（eg: BTC_USDT） |
| start_time | 结束时间，以秒计数的时间戳，不限为0 | string | 时间戳（s）            |
| end_time   | 结束时间，以秒计数的时间戳，不限为0 | string | 时间戳（s）            |
| offset     | 偏移                                | string | 自定义                 |
| limit      | 限制                                | string | 不超过100              |
| side       | 1为ASK卖出，2为BID买入,不限为0      | string | 0，1，2                |

### 示例
```
Request: 
POST https://www.ztb.im/api/v1/private/order/finished

Response:
{
  "code": 0,
  "message": "操作成功",
  "result": {
    "limit": 2,
    "offset": 0,
    "records": [
      {
        "amount": "1",
        "ctime": 1535538409.189721,
        "deal_fee": "0.00019607843",
        "deal_money": "0.999999993",
        "deal_stock": "0.19607843",
        "ftime": 1535538409.189735,
        "id": 32868,
        "maker_fee": "0",
        "market": "BTC_USDT",
        "price": "0",
        "side": 2,
        "source": "web,1",
        "taker_fee": "0.001",
        "type": 2,
        "user": 670865
      },
      {
        "amount": "10",
        "ctime": 1535538403.233823,
        "deal_fee": "0.001109999955",
        "deal_money": "1.109999955",
        "deal_stock": "0.21764705",
        "ftime": 1535538409.189735,
        "id": 32867,
        "maker_fee": "0.001",
        "market": "BTC_USDT",
        "price": "5.1",
        "side": 1,
        "source": "web,1",
        "taker_fee": "0.001",
        "type": 1,
        "user": 670865
      }
    ]
  }
}
```

### 返回数据说明

``` 
amount: 数量
ctime: 创建时间
deal_fee: 成交手续费
deal_money: 成交金额
deal_stock: 成交资产
id: 编号
left: 剩余
maker_fee: maker手续费
market: 市场名
ftime: 完成时间
price: 价格
side: 1为ASK卖出，2为BID买入
source:来源
taker_fee: taker手续费
type: 交易类型，1为限价，2为市价
user: 用户编号
```

------
## 查询某个已成交订单详情

POST /api/v1/private/order/finished/detail 

频率限制：10次/s

### 请求参数

| 参数     | 描述   | 类型   | 值               |
| -------- | ------ | ------ | ---------------- |
| order_id | 订单号 | string | 自定义（eg:1081) |

### 示例

```
Request:
POST https://www.ztb.im/api/v1/private/order/finished/detail

Response:
{
  "code": 0,
  "message": "操作成功",
  "result": {
    "amount": "10",
    "ctime": 1565681925.295415,
    "deal_fee": "0",
    "deal_money": "19.5",
    "deal_stock": "10",
    "ftime": 1565681925.295421,
    "id": 1081,
    "maker_fee": "0",
    "market": "BTC_USDT",
    "price": "2",
    "side": 2,
    "source": "web,127",
    "taker_fee": "0",
    "type": 1,
    "user": 2
  }
}
```

### 返回数据说明

```
amount: 数量
ctime: 创建时间
deal_fee: 成交手续费
deal_money: 成交金额
deal_stock: 成交资产
id: 编号
left: 剩余
maker_fee: maker手续费
market: 市场名
ftime: 完成时间
price: 价格
side: 1为ASK卖出，2为BID买入
source:来源
taker_fee: taker手续费
type: 交易类型，1为限价，2为市价
user: 用户编号
```

------

## 查询订单成交接口(分笔交明细)

POST /api/v1/private/order/deals 

频率限制：10次/s

### 请求参数

| 参数     | 描述     | 类型   | 值                  |
| -------- | -------- | ------ | ------------------- |
| order_id | 订单编号 | string | 自定义(eg:32730)    |
| offset   | 偏移     | string | 自定义              |
| limit    | 限制值   | string | 自定义（不超过100） |

### 示例

```
# Request 
POST https://www.ztb.im/api/v1/private/order/deals
# Response
{
  "code": 0,
  "message": "操作成功",
  "result": {
    "limit": 20,
    "offset": 0,
    "records": [
      {
        "amount": "1",
        "deal": "19.96",
        "deal_order_id": 32730,
        "fee": "0.001",
        "id": 25503,
        "price": "19.96",
        "role": 2,
        "time": 1535437951.751402,
        "user": 670865
      }
      ...
    ]
  }
}
```

### 返回数据说明

```
limit: 限制
offset: 偏移
records: 记录
amount: 数量
deal: 已成交
deal_order_id: 成交的订单id
fee: 手续费
id: 成交id
price: 价格
role: 角色，1为Maker,2为Taker
time: 时间戳
user: 用户编号
```

---

## 

## Websocket API

### 简介

##### 接入URL

wss://ws.ztb.com/ws

##### 请求格式

- method: 请求方法，String
- params: 参数，Array
- id: 请求编号, Integer, 自定义随机数

##### 响应格式

- result: Json object，如果没有返回则为null
- error: Json object，成功返回null,如果不成功返回非null

1. code: 错误码
2. message: 错误信息

- id: 请求编号, Integer

##### 通知格式

- method: 请求方法，String
- params: 参数，Array
- id: Null

##### 通用错误码:

- 1: 参数不合法
- 2: 内部错误
- 3: 服务不可用
- 4: 方法未找到
- 5: 服务超时
- 6: 需要授权

##### 订阅主题

成功建立与Websocket服务器的连接后，Websocket客户端发送如下请求以订阅特定主题：

{"method":"method to sub","params":[request params],"id generate by client"}

```
{
  "method": "kline.subscribe",
  "params": [
    "BTC_USDT",
    300
  ],
  "id": 10086
}
```



##### 取消订阅

Websocket订阅特定主题后，如需取消订阅，Websocket客户端发送如下请求取消订阅：

{"method":"method to unsubscribe"}

```
{
  "method": "kline.unsubscribe"
}
```



##### 查询数据

Websocket服务器同时支持一次性请求数据（pull）。

请求数据的格式如下：

{"method":"method to qurey","params":[request params],"id generate by client"}

```
{
  "method": "kline.qurey",
  "params": [
    "BTC_USDT",
    1575561600,
    1575648000,
    300
  ],
  "id": 10086
}
```



## 市场状态（市场概要）

##### 主题订阅

此主题发送市场最新市场状态。

```
{"method":"state.subscribe","params":[$market$],"id":10086}
```

| 参数     | 数据类型   | 是否必需 | 描述   | 取值范围                  |
| ------ | ------ | ---- | ---- | --------------------- |
| market | string | true | 市场名称 | BTC_USDT, ETH_USDT... |

订阅请求

```
{
  "method": "state.subscribe",
  "params": [
    "BTC_USDT"
  ],
  "id": 10086
}
```

Response

```
{
  "method": "state.update",
  "params": [
    "BTC_USDT",
    {
      "last": "7461.1526",
      "volume": "15864.0388",
      "deal": "119049215.80982165",
      "period": 86400,
      "high": "7553.5791",
      "open": "7421.5379",
      "low": "7414.7222",
      "close": "7461.1526"
    }
  ],
  "id": null
}
```

##### 取消订阅

```
{"method":"state.unsubscribe"}
```

##### 查询数据

用请求方式一次性获取过去特定时间的市场状态数据.

```
{"method":"state.query","params":[$market$,$period$],"id":10086}
```

| 参数   | 数据类型 | 是否必需 | 描述     | 取值范围              |
| ------ | -------- | -------- | -------- | --------------------- |
| market | string   | true     | 市场名称 | BTC_USDT, ETH_USDT... |
| period | int      | true     | 周期     | eg: 86400             |

查询请求

```
{
  "method": "state.query",
  "params": [
    "BTC_USDT",
    86400
  ],
  "id": 10086
}
```

Response

```
{
  "error": null,
  "result": {
    "volume": "15952.03100501",
    "period": 86400,
    "deal": "119721963.749190401504",
    "last": "7467.2656",
    "open": "7431.0264",
    "low": "7428.725",
    "close": "7467.2656",
    "high": "7553.5791"
  },
  "id": 10086
}
```



## 今日市场状态（市场概要）

##### 主题订阅

此主题发送市场今日市场状态。

```
{"method":"today.subscribe","params":[$market$],"id":10086}
```

| 参数     | 数据类型   | 是否必需 | 描述   | 取值范围                  |
| ------ | ------ | ---- | ---- | --------------------- |
| market | string | true | 市场名称 | BTC_USDT, ETH_USDT... |

订阅请求

```
{
  "method": "today.subscribe",
  "params": [
    "BTC_USDT"
  ],
  "id": 10086
}
```

Response

```
{
  "method": "today.update",
  "params": [
    "BTC_USDT",
    {
      "last": "7461.1526",
      "volume": "15864.0388",
      "deal": "119049215.80982165",
      "period": 86400,
      "high": "7553.5791",
      "open": "7421.5379",
      "low": "7414.7222",
      "close": "7461.1526"
    }
  ],
  "id": null
}
```

##### 取消订阅

```
{"method":"today.unsubscribe"}
```

##### 查询数据

用请求方式一次性获取今日的市场状态数据.

```
{"method":"today.query","params":[$market$],"id":10086}
```

| 参数   | 数据类型 | 是否必需 | 描述     | 取值范围              |
| ------ | -------- | -------- | -------- | --------------------- |
| market | string   | true     | 市场名称 | BTC_USDT, ETH_USDT... |

查询请求

```
 {
  "method": "today.query",
  "params": [
    "BTC_USDT"
  ],
  "id": 10086
}
```

Response

```
{
  "error": null,
  "result": {
    "open": "7525.3908",
    "deal": "119600164.325722971504",
    "last": "7466.2622",
    "high": "7541.1691",
    "low": "7444.7897",
    "volume": "15935.75120501"
  },
  "id": 10086
}
```



## K线数据

##### 主题订阅

此主题发送最新K线数据。

```
{"method":"kline.subscribe","params":[$market$,$interval$],"id":10086}
```

| 参数       | 数据类型   | 是否必需 | 描述   | 取值范围                             |
| -------- | ------ | ---- | ---- | -------------------------------- |
| market   | string | true | 市场名称 | BTC_USDT, ETH_USDT...            |
| interval | string | true | K线周期 | 60,300,900,1800,3600,7200,14400… |

订阅请求

```
{
  "method": "kline.subscribe",
  "params": [
    "BTC_USDT", 	market
    300   			interval
  ],
  "id": 10086		id
}
```

Response

```
{
  "id": null,
  "method": "kline.update",
  "params": [
    [
      1575705900,           time
      "7542.8082",  		open
      "7534.9152", 			close
      "7547.0765", 			high
      "7530.8753", 			low
      "70.7463", 			amount
      "533428.87370982",	deal_money
      "BTC_USDT" 			market
    ]
  ]
}
```

##### 取消订阅

```
{"method":"kline.unsubscribe"}
```

#####  

##### 查询数据

用请求方式一次性获取K线数据，需要额外提供以下参数： （每次最多返回xxx条）

```
{"method":"kline.query","params":[$market$,$start$,$end$,$interval$],"id":10086}
```

| 参数    | 数据类型    | 是否必需  | 描述   |
| ----- | ------- | ----- | ---- |
| start | integer | false | 起始时间 |
| end   | integer | false | 结束时间 |

查询请求

```
{
  "method": "kline.query",
  "params": [
    "BTC_USDT",		market
    1575561600,		start
    1575648000,		end
    300				interval
  ],
  "id": 10086		id
}
```

Response

```
{
  "error": null,
  "result": [
    [
      1575561600,		time
      "7340.949",  		open
      "7345.5655", 		close
      "7357.0065", 		high
      "7332.0522", 		low
      "77.4528",		amount
      "568870.6347857", deal  成交额
      "BTC_USDT"		market
    ],
    [
      1575561900,
      "7346.2494",
      "7333.1595",
      "7350.2274",
      "7329.8995",
      "72.2238",
      "530009.12444915",
      "BTC_USDT"
    ]
    ...
}
```



## 市场深度行情数据

此主题发送最新深度行情数据。

##### 主题订阅

```
{"method":"depth.subscribe","params":[$market$,$limit$,$interval$],"id":10086}
```

| 参数       | 数据类型   | 是否必需 | 描述   | 取值范围                                     |
| -------- | ------ | ---- | ---- | ---------------------------------------- |
| market   | string | true | 市场名称 | BTC_USDT, ETH_USDT...                    |
| limite   | int    | true | 数量   | 1, 5, 10, 20, 30, 50, 100                |
| interval | string | true | 深度合并 | "0", "0.00000001", "0.0000001", "0.000001", "0.00001", "0.0001", "0.001", "0.01", "0.1" |

订阅请求

```
{
  "method": "depth.subscribe",
  "params": [
    "BTC_USDT",
    50,
    "0.0001"
  ],
  "id": 10086
}
```

Response

```
{
  "id": null,
  "method": "depth.update",
  "params": [
    true,              true 表示完整深度列表，false 表示更新
    {
      "bids": [
        [
          "7457.1469"   price
          "0.0026"    	amount
        ],
        [
          "7457.137",
          "0.0028"
        ],
        ...
      ],
      "asks": [
        [
          "7550.6256",
          "0.2271"
        ],
        [
          "7550.9482",
          "0.0022"
        ],
        ...
      ]
    },
    "BTC_USDT"
  ]
}
```

##### 取消订阅

```
{"method":"depth.unsubscribe"}
```

#####  

##### 查询数据

用请求方式一次性获取深度数据

```
{"method":"depth.query","params":[$market$,$limit$,$interval$],"id":10086}
```

| 参数       | 数据类型   | 是否必需 | 描述   | 取值范围                                     |
| -------- | ------ | ---- | ---- | ---------------------------------------- |
| market   | string | true | 市场名称 | BTC_USDT, ETH_USDT...                    |
| limite   | int    | true | 数量   | 1, 5, 10, 20, 30, 50, 100                |
| interval | string | true | 深度合并 | "0", "0.00000001", "0.0000001", "0.000001", "0.00001", "0.0001", "0.001", "0.01", "0.1" |

查询请求

```
{
  "method": "depth.query",
  "params": [
    "BTC_USDT",
    10,
    "0.0001"
  ],
  "id": 10086
}
```

Response

```
{
  "id": 10086,
  "error": null,
  "result": {
    "asks": [
      [
        "7562.2075",
        "0.0228"
      ],
      [
        "7577.5392",
        "0.001"
      ],
      ...
    ],
    "bids": [
      [
        "7477.9723",
        "0.2047"
      ],
      [
        "7477.9225",
        "0.3294"
      ],
      ...
    ]
  }
}
```



## 市场最新价格数据

此主题发送市场最新价格。

##### 主题订阅

```
{"method":"price.subscribe","params":[$market$],"id":10086}
```

| 参数   | 数据类型 | 是否必需 | 描述     | 取值范围              |
| ------ | -------- | -------- | -------- | --------------------- |
| market | string   | true     | 市场名称 | BTC_USDT, ETH_USDT... |

订阅请求

```
{
  "method": "price.subscribe",
  "params": [
    "BTC_USDT",
  ],
  "id": 10086
}
```

Response

```
{
  "method": "price.update",
  "params": [
    "BTC_USDT",
    "7514.2520"
  ],
  "id": null
}
```

##### 取消订阅

```
{"method":"price.unsubscribe"}
```

#####  

##### 查询数据

用请求方式一次性获取市场最新价格数据

```
{"method":"price.query","params":[$market$],"id":10086}
```

| 参数   | 数据类型 | 是否必需 | 描述     | 取值范围               |
| ------ | -------- | -------- | -------- | ---------------------- |
| market | string   | true     | 市场名称 | BTC_USDT, ETH_USDT,... |

查询请求

```
{
  "method": "price.query",
  "params": [
    "BTC_USDT"
  ],
  "id": 10086
}
```

Response

```
{
  "error": null,
  "result": "7482.0109",
  "id": 10086
}
```



## 最新成交数据

此主题发送市场最新成交数据。

##### 主题订阅

```
{"method":"deals.subscribe","params":[$market$],"id":10086}
```

| 参数   | 数据类型 | 是否必需 | 描述     | 取值范围               |
| ------ | -------- | -------- | -------- | ---------------------- |
| market | string   | true     | 市场名称 | BTC_USDT, ETH_USDT,... |

订阅请求

```
{
  "method": "deals.subscribe",
  "params": [
    "BTC_USDT",
  ],
  "id": 10086
}
```

Response

```
{
  "method": "deals.update",
  "params": [
    "BTC_USDT",
    [
      {
        "id": 597933730,
        "time": 1575876545.1941223,
        "type": "sell",
        "price": "7477.6154",
        "amount": "0.1416"
      }
    ]
  ],
  "id": null
}
```

##### 取消订阅

```
{"method":"deals.unsubscribe"}
```

#####  

##### 查询数据

用请求方式一次性获取市场最新成交数据。

```
{"method":"deals.query","params":[$market$,$limit$,$last_id$],"id":10086}
```

| 参数      | 数据类型   | 是否必需 | 描述          | 取值范围                      |
| ------- | ------ | ---- | ----------- | ------------------------- |
| market  | string | true | 市场名称        | BTC_USDT, ETH_USDT ...    |
| limit   | int    | true | 数量          | 1, 5, 10, 20, 30, 50, 100 |
| last_id | int    | true | 上次返回结果的最大id | 597967944                 |

查询请求

```
{
  "method": "deals.query",
  "params": [
    "BTC_USDT",
    10,
    598129296
  ],
  "id": 10086
}
```

Response

```
{
  "error": null,
  "result": [
    {
      "id": 598136190,
      "type": "sell",
      "time": 1575881302.0342646,
      "price": "7459.6875",
      "amount": "0.1781"
    },
    {
      "id": 598136185,
      "type": "sell",
      "time": 1575881301.876456,
      "price": "7463.8087",
      "amount": "0.2333"
    },
    ...s
  ],
  "id": 10086
}
```

