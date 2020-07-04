package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

func writeMessage(conn *websocket.Conn, message []byte) {
	err := conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		fmt.Println("write error", err)
		conn.Close()
		os.Exit(0)
	}
}

func readMessage(c *websocket.Conn) {
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			log.Printf("ws-read-err: %s\n", err)
			c.Close()
			break
		}
		log.Println(string(message))
	}
}

func main() {
	// 这个地址是客户端订阅pushservice的地址
	// wss://futurews.zgnext.com/realtime 测试
	// wss://futurews.51meeting.com/realtime 线上

	// new WebSocket("ws://localhost:3000?token=xxxxxxxxxxxxxxxxxxxx");
	// 在请求头中加入token
	// 大体上Websocket的身份认证都是发生在握手阶段，通过请求中的内容来认证
	// 后端要解析token,
	// 如果从token中解析到用户信息和请求中的用户信息一致，则是登录的用户，否则校验不通过
	//token := "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiI2MjA0NzQyNDgwNzg5MjcxMDQiLCJfdGltZSI6MTU5MjE5MzQyMTgzOCwiX3IiOiJoM2p6MEpvMXJiRU8iLCJfcCI6IjE3MGU0NWFhOWY0YjUwZjU1NDI5YmFhM2Q4MjIwNjg5In0.OfAqj3UlFheyrT4spgxLjoQpdQyHOD8Gz8Us2h9xNPI"
	//requestHeader := http.Header{}
	//requestHeader.Set("token", token)
	//https://ss.abkjl.com/v1/rest/GetLatestKLine
	// wss://futurews.zg8.com/realtime （线上）
	// wss://futurews.51meeting.com/realtime （线上备用）
	c, resp, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8187/realtime", nil)
	//c, resp, err := websocket.DefaultDialer.Dial("wss://futurews.51meeting.com/realtime", nil)
	//c, resp, err := websocket.DefaultDialer.Dial("wss://futurews.zg8.com/realtime", nil)
	//c, resp, err := websocket.DefaultDialer.Dial("wss://futurews.zgnext.com/realtime", nil)
	//c, resp, err := websocket.DefaultDialer.Dial("wss://api-aws.huobi.pro/ws", nil)
	//c, resp, err := websocket.DefaultDialer.Dial("wss://ss.abkjl.com/v1/market", nil)
	if err != nil {
		if resp != nil {
			log.Println("status", resp.StatusCode)
			if resp.Body != nil {
				respBody, _ := ioutil.ReadAll(resp.Body)
				log.Println("info", string(respBody))
			}
		}
		log.Fatal("dial:", err)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("websocket connected", resp.StatusCode, string(respBody))

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		readMessage(c)
	}()

	//subkline := `{"rid":"4","req":"Sub","args":["index_BTC", "kline_1m_BTC", "kline_3m_BTC", "kline_5m_BTC",
	//"kline_15m_BTC", "kline_30m_BTC", "kline_1h_BTC", "kline_2h_BTC", "kline_4h_BTC", "kline_6h_BTC",
	//"kline_12h_BTC", "kline_1d_BTC", "kline_7d_BTC", "kline_30d_BTC"]}`

	//subkline := `{"rid":"4","req":"Sub","args":["kline_1m_BTC", "__slow__"]}`

	//subquotation := `{"rid":"4","req":"Sub","args":["index_BTC"]}`
	//subkline := `{"rid":"4","req":"Sub","args":["kline_1m_BTC"]}`
	//subPosition := `{"rid":"4","req":"Sub","args":["position_TESTUSD"]}`
	//subPing := `ping`
	//message := []byte("{\"sub\":\"market.btcusdt.kline.1min\", \"id\": \"id1\"}")

	// 盖亚的订阅
	// rid:10 kline_1m
	subkline := `{"req":"Sub","rid":"20","args":["kline_1w_BTC.USDT"],"expires":1593766170283}`
	//{"req":"Sub","rid":"10","args":["kline_1m_BTC.USDT"],"expires":1593766170283}
	//subquotation := `{"rid":"88","req":"Sub","args":["index_GMEX_CI_BTC","__slow__"],"expires":"1588848662293"}`
	//subHbHisKline := `{"req": "market.btcusdt.kline.1min", "id": "id1", "from": 1592792400, "to": 1592792700}`

	//history := `{"req":"GetHistKLine","rid":"103","args":{"Sym":"BTC","Typ":"1w","Offset":0,"Sec":1588416631,"Count":1},"expires":"1588848662118"}`
	//history := `{"req":"GetHistKLine","rid":"103","args":{"Sym":"GMEX_CI_BTC","Typ":"1m","Offset":0,"Sec":1588416631,"Count":1},"expires":"1588848662118"}`
	//subquotation := `{"rid":"88","req":"Sub","args":["index_GMEX_CI_BTC","__slow__"],"expires":"1588848662293"}`
	//unsubKline := `{"rid":"4","req":"UnSub","args":["kline_1m_BTC"],"expires":"1588848658882"}`
	//unsubIndex := `{"rid":"4","req":"UnSub","args":["index_BTC"],"expires":"1588848658882"}`

	//subTime := `{"req":"Time", "rid":"6", "expires": 1593340347000, "args":1593340287707}`

	//{"args":["index_BTC"],"expires":1593602105790,"req":"UnSub","event":"cancel","id":"10048","key":"index","rid":"20","sendTime":0}

	//writeMessage(c, []byte(history))
	writeMessage(c, []byte(subkline))
	//writeMessage(c, []byte(subquotation))
	//writeMessage(c, []byte(subPosition))
	//writeMessage(c, []byte(subTime))

	wg.Wait()
}
