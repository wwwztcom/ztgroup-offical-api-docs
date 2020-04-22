# ZT OPEN API

[简体中文](./readme.md) |English

- [**GUIDE**](#Guide)
  
  - [**Create API key**](#Create API key)
  - [**Interface call mode description**](#Interface call mode description)
- [**Server**](#The Server)
  
- [**REST API**](#rest-api)
  
  - [**Access URL**](#接入-url)
  - [**Request  interaction**](#Request  interaction)
  - [**！！Be careful！！**](#Caution)
  - [**Get exchange market**](#Get exchange market)
  - [**Get  depth information**](#Get  depth information)
  - [**Get the latest transaction record**](#Get the latest transaction record)
  - [**Get K-Line Information**](#Get K-Line Information)
  - [**Get symbol information**](#Get symbol information)
  - [**Signature authentication**](#Signature authentication)
  - [**User assets**](#Get user assets)
  - [**Limit order**](#Limit order)
  - [**Market order**](#Market order)
  - [**Cancel an order**](#Cancel an order)
  - [**Cancel orders in batch**](#**Cancel orders in batch**)
  - [**Get pending orders**](#Get pending orders)
  - [**Get a pending order detail**](#Get a pending order detail)
  - [**Get finished orders**](#Get finished orders)
  - [**Get a finished order detail **](#Get a finished order detail )
  
- [**Websocket API**](#websocket-api)

  - [**Introduce**](#Introduce) 
  - [**Market status**](#Market status)
  - [**Market status today**](#Market status today)
  - [**K-Line Data**](#K-Line Data)
  - [**Market depth**](#Market depth)
  - [**Latest price**](#Latest price)
  - [**Latest trade history**](#Latest trade history)

  

## Guide

**This is the developer documentation, ZT provides a simple and easy-to-use API interface, through which you can obtain market data, conduct transactions, and manage orders**

### Create API key

After registering an account with **[ZT](https://www.zt.com)** , the user needs to create an API key in [API management]. After the creation, a set of randomly generated API key and secret key will be obtained. With this set of data, programmed transactions can be carried out. A single account can create up to 5 keys

> **Please do not disclose the information of API key and secret key, so as to avoid asset loss. It is recommended that the user bind the IP address for the API. Each key is bound to a maximum of 5 IPS, separated by English commas**

### Interface call mode description

- REST API

  It provides functions of market query, balance query, currency transaction and order management. It is recommended that users use rest API to query account balance, currency transaction and order management

### The Server

- ZT server runs in Tokyo. In order to minimize the API access delay, it is recommended to use a server with smooth communication with Tokyo


## REST API

### Access URL

- [https://www.zt.com](https://www.zt.com) 

### Request  interaction

#### Introdue

Rest API provides functions of market inquiry, balance inquiry, currency transaction and order management

All requests are based on the HTTPS protocol. The content type in the request header information needs to be set to the form format:

- **content-type:application/x-www-form-urlencoded**

#### Error code 

| **Error code** | **introduce**                                | **reason**               |
| :------------- | :------------------------------------------- | :----------------------- |
| 0              | Success                                      |                          |
| 1              | arameter is invalid                          |                          |
| 2              | Internal error                               |                          |
| 3              | ervice is not available                      |                          |
| 4              | Method not found                             |                          |
| 5              | Service timeout                              |                          |
| 10             | Insufficient amount                          |                          |
| 11             | The number of transactions is too small      |                          |
| 12             | Insufficient depth                           |                          |
| 10005          | Record not found                             | X-SITE-ID，Setting error |
| 10022          | User does not have a real name               |                          |
| 10051          | User prohibited transaction                  |                          |
| 10056          | Less than minimum amoun                      |                          |
| 10059          | No transaction has been opened for the asset |                          |
| 10060          | The transaction has not been opened          |                          |
| 10062          | Incorrect amount accuracy                    |                          |

### Caution

- All interface requests (public and private) must add an x-site-id field to the header of the request request, with a value of "1". This field is not used for signature verification.

# Market API public interface

## Get exchange market

Get  /api/v1/tickers  

Frequency limit：20times/s

### example

```
Request:
GET https://www.zt.com/api/v1/tickers

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

---

## Get  depth information

Get /api/v1/depth 

Frequency limit: 20 times / S

### Request parameters

| parameters | **description**           |
| ---------- | ------------------------- |
| symbol     | Market name               |
| size       | Number of  returned depth |

### Example

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

### **Description of return data**
```
asks : ask depth[price，amount]
bids : bid depth[price，amount]
```
---

## **Get the latest transaction record**

Get /api/v1/trades  

Frequency limit: 20 times / S

### Request parameters

| parameters | **description**           |
| ---------- | ------------------------- |
| symbol     | Market name               |
| size       | Number of returned record |

### Example

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

---

## Get K-Line Information

Get /api/v1/kline  

Frequency limit: 20 times / S

### Request parameters

| parameters | **description**                                              |
| ---------- | ------------------------------------------------------------ |
| symbol     | Market name                                                  |
| type       | Time sharing parameter, which can be 1min,5min,15min,30min,hour,day,week |
| size       | Number of returned transaction                               |

### Example

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



---

## Get symbol information

Get /api/v1/exchangeInfo

Frequency limit: 20 times / S

### Example

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



---
# Trading API Private Interface

- All private interface requests use the POST method, and parameters are submitted in the form of data (all private interface api_key and sign must be passed).

## Signature authentication

- All parameters must be verified by signature, and all parameters must be sorted by parameter name according to alphabet.

### Example

```
Request parameters：

'market':'BTC_USDT',
'side':1,
'price':'50',
'amount':'0.02'

Sorted parameter string:amount=0.02&api_key=apiKey&market=BTC_USDT&price=50&side=1

Note: to generate MD5 signature, secret_key must be added to the string generated above to generate the final string

Final signature string:
amount=0.02&api_key=apiKey&market=BTC_USDT&price=50&side=1&secret_key=secret_key

MD5 signature：
Use MD5 encryption string of 32bit, the generated encryption string must be uppercase
```

---
## Get user assets

POST /api/v1/private/user 

Frequency limit: 20 times / S

### Example

```
Request:
POST https://www.zt.com/api/v1/private/user

Response:
{
  "code": 0,
  "message": "操作成功",  Operation successful
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



---
## Limited order

POST /api/v1/private/trade/limit  

Frequency limit: 20 times / s

### Request parameters

| **parameters** | **describe**                      | **type** | **value**                  |
| -------------- | --------------------------------- | -------- | -------------------------- |
| market         | market                            | string   | customized（eg: BTC_USDT） |
| side           | 1 is ask to sell, 2 is bid to buy | string   | 1,2                        |
| amount         | amount                            | string   | customized                 |
| price          | price                             | string   | customized                 |

### Example
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



---
## Market order

POST /api/v1/private/trade/market  

Frequency limit: 20 times / S

### Request parameters

| **parameters** | **escription**                    | **type** | **value**                  |
| -------------- | --------------------------------- | -------- | -------------------------- |
| market         | market                            | string   | customized（eg: BTC_USDT） |
| side           | 1 is ask to sell, 2 is bid to buy | string   | 1,2                        |
| amount         | amount                            | string   | customized                 |

### Example

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



---
## Cancel an order

POST /api/v1/private/trade/cancel 

Frequency limit: 20 times / S

### Request parameters
| parameters | description | type   | alue                       |
| ---------- | ----------- | ------ | -------------------------- |
| market     | market      | string | customized（eg: BTC_USDT） |
| order_id   | order_id    | string | customized（eg: 32865）    |

### Example

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



---
## Cancel orders in batch

POST /api/v1/private/trade/cancel_batch 

Note : No more than 10 orders each time.

Frequency limit: 20 times / S

### Request parameters

| parameters  | Description | ype  | Value                                                        |
| ----------- | ----------- | ---- | ------------------------------------------------------------ |
| orders_json | orders json | json | eg: [{"market":"BTC_USDT", "order_id":456647},{"market":"BTC_USDT", "order_id":456648}] |

### Example

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



------
## Get pending orders

POST /api/v1/private/order/pending  

Frequency limit: 20 times / S

### Request parameters

| parameters | description | type   |                                |
| ---------- | ----------- | ------ | ------------------------------ |
| market     | market      | string | customized（eg: BTC_USDT）     |
| offset     | offset      | string | customized                     |
| limit      | limit       | string | customized（no more than 100） |


### Example
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



------
## Get a pending order detail

POST /api/v1/private/order/pending/detail  

Frequency limit: 20 times / S

### Request parameters

| **parameters** | description | type   | value                      |
| -------------- | ----------- | ------ | -------------------------- |
| market         | market      | string | customized（eg: BTC_USDT） |
| order_id       | order id    | string | customized（eg:  1080）    |

### Example
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



---
## Get finished orders

POST /api/v1/private/order/finished 

Frequency limit: 20 times / S

### Request parameters

| parameters | description                                  | type   | value                      |
| ---------- | -------------------------------------------- | ------ | -------------------------- |
| market     |                                              | string | customized（eg: BTC_USDT） |
| start_time | time stamp in seconds, unlimited to 0        | string | timestamp（s）             |
| end_time   | time stamp in seconds, unlimited to 0        | string | timestamp（s）             |
| offset     |                                              | string | customized                 |
| limit      |                                              | string | no more than 100           |
| side       | 1 is ask sell, 2 is  bid buy, unlimited to 0 | string | 0，1，2                    |

### Example
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



------
## Get a finished order detail 

POST /api/v1/private/order/finished/detail 

Frequency limit: 20 times / S

### Request parameters

| parameters | escription | type       | value            |
| ---------- | ---------- | ---------- | ---------------- |
| order_id   |            | customized | 自定义（eg:1081) |

### Example

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



---

## Websocket API

### Introduce

##### **Access URL**

wss://ws.zt.com/ws

##### Request format

- method:  String
- params: Array
- id:  Integer, customed random number

##### Response format 

- result: Json object，If not returned; otherwise null
- error: Json object，Successfully returned null, If unsuccessful, return No null

1. code: error code 
2. message: error message 

- id: Integer

##### Notice format 

- method:  String
- params:  Array
- id: Null

##### General error code:

- 1: illegal parameter
- 2: internal error
- 3: service not available
- 4:  method not found
- 5: service timeout
- 6: authorization required

##### Subscribe topics

After successfully establishing a connection with the websocket server, the websocket client sends the following request to subscribe to a specific topic:

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



##### Unsubscribe

After the websocket subscribes to a specific topic, if you need to unsubscribe, the websocket client sends the following request to unsubscribe:

{"method":"method to unsubscribe"}

```
{
  "method": "kline.unsubscribe"
}
```



##### Query 

The websocket server also supports pull.

Request format is as follows：

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



## Market status 

##### Subscribe

This topic sends the latest market status of the market.

```
{"method":"state.subscribe","params":[$market$],"id":10086}
```

| Parameter | Type   | **Required or not** | Description | **Value**    |
| --------- | ------ | ------------------- | ----------- | ------------ |
| market    | string | true                |             | eg: BTC_USDT |

Subscription request

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

##### Unsubscribe

```
{"method":"state.unsubscribe"}
```

##### Query

Obtain the market status data of the past specific time in one time by request

```
{"method":"state.query","params":[$market$,$period$],"id":10086}
```

| Parameter | Type   | Required or not | Description | **Value**   |
| --------- | ------ | --------------- | ----------- | ----------- |
| market    | string | true            |             | eg:BTC_USDT |
| period    | int    | true            |             | Eg: 86400   |

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



## Market status today

##### Subscribe

This topic is sent to market today.。

```
{"method":"today.subscribe","params":[$market$],"id":10086}
```

| Parameter | Type   | **Required or not** | Description | **Value**   |
| --------- | ------ | ------------------- | ----------- | ----------- |
| market    | string | true                |             | eg:BTC_USDT |

Subscribe request

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

##### Unsubscribe

```
{"method":"today.unsubscribe"}
```

##### Query

Obtain today's market status data at one time by request.

```
{"method":"today.query","params":[$market$],"id":10086}
```

| Parameter | Type   | **Required or not** | Description | **Value**   |
| --------- | ------ | ------------------- | ----------- | ----------- |
| market    | string | true                |             | eg:BTC_USDT |

Query request

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



## K-Line Data

##### Subscribe

This topic sends the latest K-line data.

```
{"method":"kline.subscribe","params":[$market$,$interval$],"id":10086}
```

| Parameter | Type   | **Required or not** | Description | **Value**                        |
| --------- | ------ | ------------------- | ----------- | -------------------------------- |
| market    | string | true                |             | BTC_USDT, ETH_USDT...            |
| interval  | string | true                |             | 60,300,900,1800,3600,7200,14400… |

Subscribe request

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

##### Unsubscribe

```
{"method":"kline.unsubscribe"}
```

#####  

##### Query

obtain K-line data in one time by request, the following additional parameters 

```
{"method":"kline.query","params":[$market$,$start$,$end$,$interval$],"id":10086}
```

| Parameter | Type    | Required or not | Description | **Value**      |
| --------- | ------- | --------------- | ----------- | -------------- |
| start     | integer | false           |             | eg: 1575561600 |
| end       | integer | false           |             | eg: 1575648000 |

Query request

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
      "568870.6347857", deal money
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



## Market depth

##### Subscribe

This topic sends the latest in-depth market data.

```
{"method":"depth.subscribe","params":[$market$,$limit$,$interval$],"id":10086}
```

| Parameter | Type   | Required or not | Description | **Value**                                                    |
| --------- | ------ | --------------- | ----------- | ------------------------------------------------------------ |
| market    | string | true            |             | eg: BTC_USDT                                                 |
| limite    | int    | true            |             | 1, 5, 10, 20, 30, 50, 100                                    |
| interval  | string | true            | merge depth | "0", "0.00000001", "0.0000001", "0.000001", "0.00001", "0.0001", "0.001", "0.01", "0.1" |

Subscribe request

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
    true,              true represents the full depth list，false represents update
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

##### Unsubscribe

```
{"method":"depth.unsubscribe"}
```

#####  

##### Query

Obtain depth in one time by request

```
{"method":"depth.query","params":[$market$,$limit$,$interval$],"id":10086}
```

| Parameter | Type   | Required or not | Description | Value                                                        |
| --------- | ------ | --------------- | ----------- | ------------------------------------------------------------ |
| market    | string | true            |             | eg: BTC_USDT                                                 |
| limite    | int    | true            |             | 1, 5, 10, 20, 30, 50, 100                                    |
| interval  | string | true            | merge depth | "0", "0.00000001", "0.0000001", "0.000001", "0.00001", "0.0001", "0.001", "0.01", "0.1" |

Query request

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



## Latest price 

##### Subscribe

This topic sends the latest price.

```
{"method":"price.subscribe","params":[$market$],"id":10086}
```

| Parameter | Type   | Required or not | Description | Value        |
| --------- | ------ | --------------- | ----------- | ------------ |
| market    | string | true            |             | eg: BTC_USDT |

Subscribe request

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

##### Unsubscribe

```
{"method":"price.unsubscribe"}
```

#####  

##### Query 

Get the latest price in one time by request

```
{"method":"price.query","params":[$market$],"id":10086}
```

| Parameter | Type   | Required or not | Description | Value        |
| --------- | ------ | --------------- | ----------- | ------------ |
| market    | string | true            |             | eg: BTC_USDT |

Query request

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



## Latest trade history

##### Subcribe

This topic sends the latest finished order.

```
{"method":"deals.subscribe","params":[$market$],"id":10086}
```

| Parameter | Type   | Required or not | Description | Value        |
| --------- | ------ | --------------- | ----------- | ------------ |
| market    | string | true            |             | eg: BTC_USDT |

Subcribe request

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

##### Unsubcribe

```
{"method":"deals.unsubscribe"}
```

#####  

##### Query

Obtain Latest finished order in one time by request.

```
{"method":"deals.query","params":[$market$,$limit$,$last_id$],"id":10086}
```

| Parameter | Type   | Required or not | Description                        | Value                     |
| --------- | ------ | --------------- | ---------------------------------- | ------------------------- |
| market    | string | true            |                                    | eg: BTC_USDT,             |
| limit     | int    | true            |                                    | 1, 5, 10, 20, 30, 50, 100 |
| last_id   | int    | true            | Maximum ID of last returned result | eg: 597967944             |

Query request

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

