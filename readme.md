# ZT OPEN API


- [**入门指引**](#入门指引)
  - [**创建API Key**](#创建api-key)
  - [**接口调用方式说明**](#接口调用方式说明)
  - [**服务器**](#服务器)

- [**REST API**](#rest-api)
  - [**接入 URL**](#接入-url)
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
  - [**查询订单成交接口**](#查询订单成交接口)
  - [**查询未成交订单**](#查询未成交订单)
  - [**查询某个未成交订单详情**](#查询某个未成交订单详情)
  - [**查询已成交订单接口**](#查询已成交订单接口)
  - [**查询某个已成交订单详情**](#查询某个已成交订单详情)

## 入门指引

**欢迎使用开发者文档，ZT提供了简单易用的API接口，通过API可以获取市场行情数据、进行交易、管理订单**

### 创建API Key

用户在 **[ZT](https://www.zt.com)** 注册账号后，需要在[API管理]中创建API Key秘钥，创建完成后得到一组随机生成的API Key与Secret Key,利用这一组数据可以进行程序化交易，单个账号最多创建5个密钥

> **请不要泄露API Key 与 Secret Key信息，以免造成资产损失,建议用户为API绑定IP地址，每个密钥最多绑定5个IP，使用英文逗号进行分隔**

### 接口调用方式说明

- REST API

  提供行情查询、余额查询、币币交易、订单管理功能，建议用户使用REST API进行账户余额查询、币币交易及订单管理等操作

### 服务器

- ZT服务器运行在东京，为了最大限度地减少API访问延迟，建议使用与东京通讯通畅的服务器


## REST API

### 接入 URL

- [https://www.zt.com](https://www.zt.com) 

### 注意

- 所有接口请求（公共和私有接口）都必须在Request请求的Header中添加 X-SITE-ID 字段，该字段的值为”1“。该字段不用做签名校验。


# 行情API 公共接口

## 获取交易所市场数据

Get  /api/v1/tickers  

频率限制：20次/s

### 示例

```
Request:
GET https://www.zt.com/api/v1/tickers

Response:
{
    "ticker":[
        {
            "buy":"0.378",
            "high":"0.39999995",
            "last":"0.388",
            "low":"0.374101",
            "sell":"0.387",
            "symbol":"BTC_USDT",
            "vol":"3485328.1114718"
        },
        {
            "buy":"1924",
            "high":"1938.84",
            "last":"1924",
            "low":"1864.97",
            "sell":"1926",
            "symbol":"ETH_USDT",
            "vol":"2948.19477569"
        }
    ],
    "timestamp":"1535452275851"
}
```

### 返回数据说明
```
timestamp: 服务端时间戳
buy: 最佳BID
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
GET https://www.zt.com/api/v1/depth?symbol=BTC_USDT&size=1

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
GET https://www.zt.com/api/v1/trades?symbol=BTC_USDT&size=1

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
GET https://www.zt.com/api/v1/kline?symbol=BTC_USDT&type=1min&size=10

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

Get /api/v1/trades  

频率限制：20次/s

### 示例
```
Request:
GET https://www.zt.com/api/v1/exchangeInfo

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

- 所有私有接口请求都使用POST方法，参数以form-data的形式提交

## 签名认证

- 所有的参数都必须进行签名验证，所有参数必须根据字母表按照参数名进行排序

### 示例

```
请求参数：

'market':'BTC_USDT',
'side':1,
'price':'50',
'amount':'0.02'

参数字符串:amount=33.33&api_key=apiKey&market=BTC_USDT&price=50&side=1

注意:生成MD5签名必须要secretKey,在以上生成的字符串基础上添加secret_key以生成最终的字符串。

最终签名字符串:amount=0.02&api_key=apiKey&market=BTC_USDT&price=50&side=1&secret_key=secretKey

MD5签名：
使用32bit的MD5加密字符串，生成的加密字符串必须大写
```

---
## 获取用户资产

POST /api/v1/private/user 

频率限制：500次/min

### 示例

```
Request:
POST https://www.zt.com/api/v1/private/user

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

频率限制：500次/min

### 请求参数

| 参数     | 描述              |
| ------ | --------------- |
| market | 市场              |
| side   | 1为ASK卖出，2为BID买入 |
| amount | 数量              |
| price  | 价格              |

### 示例
```
Request:
POST https://www.zt.com/api/v1/private/trade/limit

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

频率限制：500次/min

### 请求参数

| 参数     | 描述              |
| ------ | --------------- |
| market | 市场              |
| side   | 1为ASK卖出，2为BID买入 |
| amount | 数量              |

### 示例

```
Request: 
POST https://www.zt.com/api/v1/private/trade/market

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

频率限制：500次/min

### 请求参数
| 参数       | 描述   |
| -------- | ---- |
| market   | 市场名称 |
| order_id | 订单编号 |

### 示例

```
Request: 
POST https://www.zt.com/api/v1/private/trade/cancel

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
maker_fee: maker手续费
market: 市场名
mtime: 发布到市场时间
price: 价格
side: 1为ASK卖出，2为BID买入
source:来源
taker_fee: taker手续费
type: 交易类型，1为限价，2为市价
user: 用户编号
```

---
## 批量取消委托订单

POST /api/v1/private/trade/cancel_batch 每次批量取消委托订单数量不超过10个。

频率限制：500次/min

### 请求参数

| 参数         | 描述      |
| ---------- | ------- |
| order_json | 订单编号    |
| sign       | 签名      |
| api_key    | api_key |

### 示例

```
Request:
POST https://www.zt.com/api/v1/private/trade/cancel_batch

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
## 查询订单成交接口

POST /api/v1/private/order/deals 

频率限制：500次/min

### 请求参数

| 参数       | 描述   |
| -------- | ---- |
| order_id | 订单编号 |
| offset   | 偏移   |
| limit    | 限制值  |

### 示例
```
# Request 
POST https://www.zt.com/api/v1/private/order/deals
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
## 查询未成交订单

POST /api/v1/private/order/pending  

频率限制：500次/min

### 请求参数

| 参数     | 描述   |
| ------ | ---- |
| market | 市场   |
| offset | 偏移   |
| limit  | 限制值  |


### 示例
```
# Request 
POST https://www.zt.com/api/v1/private/order/pending
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
taker_fee: taker手续费
type: 交易类型，1为限价，2为市价
user: 用户编号
```

------
## 查询某个未成交订单详情

POST /api/v1/private/order/pending/detail  

频率限制：500次/min

### 请求参数

| 参数       | 描述   |
| -------- | ---- |
| market   | 市场   |
| Order_id | 订单号  |

### 示例
```
Request:
POST https://www.zt.com/api/v1/private/order/pending/detail

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
maker_fee: maker手续费
market: 市场名
ftime: 发布到市场时间
price: 价格
side: 1为ASK卖出,2为BID买入
source: 来源
taker_fee: taker手续费
type: 交易类型,1为限价,2为市价
user: 用户编号
```

---
## 查询已成交订单接口

POST /api/v1/private/order/finished 

频率限制：500次/min

### 请求参数

| 参数         | 描述                   |
| ---------- | -------------------- |
| market     | 市场                   |
| start_time | 结束时间，以秒计数的时间戳，不限为0   |
| end_time   | 结束时间，以秒计数的时间戳，不限为0   |
| offset     | 偏移                   |
| limit      | 限制                   |
| side       | 1为ASK卖出，2为BID买入,不限为0 |

### 示例
```
Request: 
POST https://www.zt.com/api/v1/private/order/finished

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

频率限制：500次/min

### 请求参数

| 参数       | 描述   |
| -------- | ---- |
| order_id | 订单号  |

### 示例

```
Request:
POST https://www.zt.com/api/v1/private/order/finished/detail

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


