# 行情API

行情api包括私有api都需要在http request请求的header中添加 X-SITE_ID 字段，该字段的值为”1“。该字段不用做签名校验。

## 获取交易所市场数据

Get /api/v1/tickers  获取价格数据

https://www.zt.com/api/v1/tickers

### 示例
```
# Request 
GET https://www.zt.com/api/v1/tickers
# Response
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

### 返回数据

    timestamp: 服务端时间戳
    buy: 最佳BID
    high: 最高价
    last: 最新价
    low: 最低价
    sell: 最佳ASK
    vol: 成交量 (24小时交易量)

---
## 深度接口

Get /api/v1/depth  获取深度数据

https://www.zt.com/api/v1/depth

### 示例
```
# Request 
GET https://www.zt.com/api/v1/depth?symbol=BTC_USDT&size=1
# Response
{
    "asks":[
        ["0.387","43189.58824906"]
    ],
    "bids":[
        ["0.378","2078.91534391"]
    ]
}
```

### 返回数据
    asks : ask深度[价格，数量]
    bids : bid深度[价格，数量]

### 请求参数

| 参数     | 描述      |
| ------ | ------- |
| symbol | 市场名称    |
| size   | 返回深度的条数 |



---
## 最新成交记录接口

Get /api/v1/trades  获取最新成交记录数据

https://www.zt.com/api/v1/trades

### 示例
```
# Request 
GET https://www.zt.com/api/v1/trades?symbol=BTC_USDT&size=1
# Response
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

### 返回数据
    amount : 数量
    price : 价格
    side: 买入或者卖出
    timestamp: 时间戳

### 请求参数
| 参数     | 描述      |
| ------ | ------- |
| symbol | 市场名称    |
| size   | 返回深度的条数 |



---
## K线接口

Get /api/v1/trades  获取最新成交记录数据

https://www.zt.com/api/v1/kline

### 示例
```
# Request 
GET https://www.zt.com/api/v1/kline?symbol=BTC_USDT&type=1min&size=100
# Response
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
]
```

### 返回数据

```
[
    1532655300000: 时间戳 
    2370.16: 开
    2380: 高
    2352: 低
    2367.37: 收 
    17259.83: 交易量
]
```

### 请求参数
| 参数     | 描述                                       |
| ------ | ---------------------------------------- |
| symbol | 市场名称                                     |
| type   | 分时参数，可以为1min,5min,15min,30min,hour,day,week |
| size   | 返回成交数的条数                                 |

---
## 交易对信息接口

Get /api/v1/trades  获取交易对信息数据

https://www.zt.com/api/v1/exchangeInfo

### 示例
```
# Request 
GET https://www.zt.com/api/v1/exchangeInfo
# Response
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

### 返回数据

```
[
    baseAsset: 基础货币 
    baseAssetPrecision: 基础货币精度
    quoteAsset: 计价货币
    quoteAssetPrecision: 计价货币精度
    status: 交易状态
    symbol: 交易对
]
```

---
# 交易API private

所有的参数都需要使用form-data的形式提交数据，接口都为POST形式
交易api需要进行验签
除了sign参数外所有的参数都必须进行签名，所有参数必须根据字母表按照参数名进行排序
举例，如果请求的参数是
```
{'api_key':'c821db84-6fbd-11e4-a9e3-c86000d26d7c','sign':'5C263F54B613EEFE02BB596879D1DDF3','symbol':'BTC_USDT','side':1,'price':'50','amount':'0.02'}}")
```
amount=1.0&api_key=c821db84-6fbd-11e4-a9e3-c86000d26d7c&price=680&side=1&symbol=BTC_USDT
字符串为:
amount=1.0&api_key=c821db84-6fbd-11e4-a9e3-c86000d26d7c&price=680&side=1&symbol=BTC_USDT

MD5签名
生成MD5签名必须要secretKey,在以上生成的字符串后面添加secret_key以生成最终的字符串，例如amount=1.0&api_key=c821db84-6fbd-11e4-a9e3-c86000d26d7c&price=680&side=1&symbol=BTC_USDT&secret_key=secretKey
注意: '&secret_key=secretKey'必填. 使用32bit的MD5加密字符串，将生成的签名传到sign参数，生成的加密字符串必须大写

---
## 获取用户资产接口

POST /api/v1/private/user  获取用户资产数据

https://www.zt.com/api/v1/private/user

### 示例
```
# Request 
POST https://www.zt.com/api/v1/private/user
# Response
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

### 返回数据

```
[
    available: 可用余额
    freeze: 交易冻结金额
    other_freeze: 其他冻结金额（包括C2C和提币冻结）
    recharge_status: 充值状态0为不可充值，1为可充值
    withdraw_fee: 提币手续费
    withdraw_max: 最大提币金额
    withdraw_min: 最小提币金额
    withdraw_status: 提币状态0为不可提币，1为可提币
]
```

---
## 用户限价交易接口

POST /api/v1/private/trade/limit  用户限价交易

https://www.zt.com/api/v1/private/trade/limit

### 示例
```
# Request 
POST https://www.zt.com/api/v1/private/trade/limit
# Response
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

### 返回数据

```
[
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
]
```

### 请求参数
| 参数     | 描述              |
| ------ | --------------- |
| market | 市场              |
| side   | 1为ASK卖出，2为BID买入 |
| amount | 数量              |
| price  | 价格              |



---
## 用户市价交易接口

POSTs /api/v1/private/trade/market  用户市价交易

https://www.zt.com/api/v1/private/trade/market

### 示例
```
# Request 
POST https://www.zt.com/api/v1/private/trade/market
# Response
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

### 返回数据

```
[
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
]
```

### 请求参数

| 参数     | 描述              |
| ------ | --------------- |
| market | 市场              |
| side   | 1为ASK卖出，2为BID买入 |
| amount | 数量              |
|        |                 |



---
## 用户取消交易接口

POSTs /api/v1/private/trade/cancel  用户取消交易

https://www.zt.com/api/v1/private/trade/cancel

### 示例
```
# Request 
POST https://www.zt.com/api/v1/private/trade/cancel
# Response
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

### 返回数据

```
[
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
]
```

### 请求参数
| 参数       | 描述   |
| -------- | ---- |
| market   | 市场名称 |
| order_id | 订单编号 |

---
## 查询订单成交接口

POSTs /api/v1/private/order/deals  查询订单成交

https://www.zt.com/api/v1/private/order/deals

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

### 返回数据

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

### 请求参数
| 参数       | 描述   |
| -------- | ---- |
| order_id | 订单编号 |
| offset   | 偏移   |
| limit    | 限制值  |



---
## 查询用户未成交接口

POSTs /api/v1/private/order/pending  查询用户未成交

https://www.zt.com/api/v1/private/order/pending

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

### 返回数据

```
[
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
]
```

### 请求参数
| 参数     | 描述   |
| ------ | ---- |
| market | 市场   |
| offset | 偏移   |
| limit  | 限制值  |



---
## 查询用户已成交接口

POSTs /api/v1/private/order/finished  查询用户已成交

https://www.zt.com/api/v1/private/order/finished

### 示例
```
# Request 
POST https://www.zt.com/api/v1/private/order/finished
# Response
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

### 返回数据

```
[
    amount: 数量
    ctime: 创建时间
    deal_fee: 成交手续费
    deal_money: 成交金额
    deal_stock: 成交资产
    id: 编号
    left: 剩余
    maker_fee: maker手续费
    market: 市场名
    ftime: 完成时间
    price: 价格
    side: 1为ASK卖出，2为BID买入
    source:来源
    taker_fee: taker手续费
    type: 交易类型，1为限价，2为市价
    user: 用户编号
]
```

### 请求参数
| 参数         | 描述                   |
| ---------- | -------------------- |
| market     | 市场                   |
| start_time | 结束时间，以秒计数的时间戳，不限为0   |
| end_time   | 结束时间，以秒计数的时间戳，不限为0   |
| offset     | 偏移                   |
| limit      | 限制                   |
| side       | 1为ASK卖出，2为BID买入,不限为0 |



