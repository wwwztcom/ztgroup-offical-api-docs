package exchange_api_demo_golang

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/url"
	"time"
)

/*
api 请求地址 wss://ws.ztb.com/ws

请求

- method: 请求方法，String
- params: 参数，Array
- id: 请求编号, Integer

响应

- result: Json object，如果没有返回则为null
- error: Json object，成功返回null,如果不成功返回非null

1. code: 错误码
2. message: 错误信息

- id: 请求编号, Integer

通知

- method: 请求方法，String
- params: 参数，Array
- id: Null

*/

// 订阅
func SubClient() {

	u := url.URL{Scheme: "wss", Host: "ws.ztb.com", Path: "/ws"}
	dialer := websocket.DefaultDialer

	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Print("dial ws server failed ", err)
		return
	}

	defer func() {
		conn.Close()
	}()

	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(10000)

	/*
		订阅市场状态（市场概要）
		方法：
		status.subscribe

		参数：
		1、market string：市场名称, BTC_CNT
		2、period int: 周期，（如果订阅过去24小时市场状态，则填写 86400 ）

		eg: {"method":"state.subscribe","params":["BTC_USDT",86400],"id":10086}
	*/

	msg := []byte(fmt.Sprintf(`{"method":"today.subscribe","params":["BTC_USDT"],"id":%d}`, randNum))


	/*
		订阅kline信息
		方法：
		kline.subscribe

		参数：
		1、market string：市场名称, BTC_CNT
		2、interval int: 间隔（K线类型）, 单位为秒（s), 300表示 5min k线

		eg: {"method":"kline.subscribe","params":["BTC_USDT",300],"id":10086}
	*/

	/*msg := []byte(fmt.Sprintf(`{"method":"kline.subscribe","params":["BTC_USDT",300],"id":%d}`, randNum))
	 */

	/*
		订阅深度信息
		方法：
		depth.subscribe

		参数：
		1、market string: 市场名称, BTC_CNT
		2、limit int: 数量限制，1, 5, 10, 20, 30, 50, 100
		3、interval string: 深度合并，合并的规则为 "0", "0.00000001", "0.0000001", "0.000001", "0.00001", "0.0001", "0.001", "0.01", "0.1"

		eg: {"method":"depth.subscribe","params":["BTC_USDT",50,"0.0001"],"id":10086}
	*/

	/*msg := []byte(fmt.Sprintf(`{"method":"depth.subscribe","params":["BTC_USDT",50,"0.0001"],"id":%d}`, randNum))
	 */

	/*
		订阅市场最新价格
		方法：
		price.subscribe

		参数：
		1、市场列表 []string : 市场列表, ["BTC_USDT","ETH_USDT"]

		eg: {"method":"price.subscribe","params":["BTC_USDT","ETH_USDT"],"id":10086}
	*/

	// msg := []byte(fmt.Sprintf(`{"method":"price.subscribe","params":["BTC_USDT","ETH_USDT"],"id":%d}`, randNum))

	/*
		订阅市场最新成交
		方法：
		deals.subscribe

		参数：
		1、market string : 市场名称 BTC_USDT
	*/

	// msg := []byte(fmt.Sprintf(`{"method":"deals.subscribe","params":["BTC_USDT"],"id":%d}`, randNum))

	logs.Info(string(msg))
	err = conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		fmt.Print("write err:", err)
		return
	}

	/*msg = []byte(`{"method":"price.unsubscribe"}`)
	err = conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		fmt.Print("write err:", err)
		return
	}*/

	// 心跳消息 15s发送一次
	go func() {

		pingTicker := time.NewTicker(time.Second * 15)

		for {
			select {
			case <-pingTicker.C:

				randNum = rand.Intn(100000)
				msg = []byte(fmt.Sprintf(`{"method":"server.ping","params":[],"id":%d}`, randNum))

				err = conn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					fmt.Print("send msg err:", err)
					return
				}
			}
		}
	}()

	// 读取响应信息
	for {
		_, readMsg, err := conn.ReadMessage()
		if err != nil {
			fmt.Print("Read Error : ", err, )
			return
		}
		logs.Info(string(readMsg))
		// fmt.Println(string(readMsg))
	}
}

// 查询
func QueryClient() {

	u := url.URL{Scheme: "wss", Host: "ws.zt.com", Path: "/ws"}
	dialer := websocket.DefaultDialer

	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Print("dial ws server failed ", err)
		return
	}

	defer func() {
		conn.Close()
	}()

	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(10000)

	/*
		查询市场状态（市场概要）
		方法：
		state.query

		参数：
		1、market string：市场名称, BTC_CNT
		2、period int: 周期，（如果订阅过去24小时市场状态，则填写 86400 ）

	eg: {"method":"state.subscribe","params":["BTC_USDT",86400],"id":10086}
	*/

	// msg := []byte(fmt.Sprintf(`{"method":"state.query","params":["BTC_USDT",86400],"id":%d}`, randNum))

	/*
		查询kline信息
		方法：
		kline.query

		参数：
		1、market string：市场名称, BTC_CNT
		2、start int: 开始时间
		3、end int: 结束时间
		4、interval int: 间隔（K线类型）, 单位为秒（s), 300表示 5min k线

		eg: {"method":"kline.query","params":["BTC_USDT",1575561600,1575648000,300],"id":10086}
	*/

	//msg := []byte(fmt.Sprintf(`{"method":"kline.query","params":["BTC_USDT",1575561600,1575648000,300],"id":%d}`, randNum))

	/*
		查询深度信息
		方法：
		depth.query

		参数：
		1、market string：市场名称, BTC_CNT
		2、limit int:  数量限制，1, 5, 10, 20, 30, 50, 100
		3、interval string: 深度合并，合并的规则为 "0", "0.00000001", "0.0000001", "0.000001", "0.00001", "0.0001", "0.001", "0.01", "0.1"

		eg: {"method":"depth.query","params":["BTC_USDT",10,0.0001],"id":10086}
	*/

	// msg := []byte(fmt.Sprintf(`{"method":"depth.query","params":["BTC_USDT",10,"0.0001"],"id":%d}`, randNum))

	/*
		    查询市场最新价格
			方法：
			price.query

			参数：
			1、market string：市场名称, BTC_CNT

			eg: {"method":"price.query","params":["BTC_USDT"],"id":10086}

	*/

	// msg := []byte(fmt.Sprintf(`{"method":"price.query","params":["BTC_USDT"],"id":%d}`, randNum))

	/*
		查询市场最新成交
		方法：
		deals.query

		参数：
		1、market string : 市场名称, BTC_USDT
		2、limit int: 数量 1, 5, 10, 20, 30, 50, 100
		3、last_id: 上次返回结果的最大id
	*/

	msg := []byte(fmt.Sprintf(`{"method":"deals.query","params":["BTC_USDT",10,598129296],"id":%d}`, randNum))

	logs.Info(string(msg))
	err = conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		fmt.Print("write err:", err)
		return
	}

	_, readMsg, err := conn.ReadMessage()
	if err != nil {
		fmt.Print("Read Error : ", err, )
		return
	}

	fmt.Println(string(readMsg))
}
