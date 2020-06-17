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

	c, resp, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8187/realtime", nil)
	//c, resp, err := websocket.DefaultDialer.Dial("wss://futurews.51meeting.com/realtime", nil)
	//c, resp, err := websocket.DefaultDialer.Dial("wss://futurews.zgnext.com/realtime", nil)
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

	//subkline := `{"rid":"4","req":"Sub","args":["kline_1m_BTC"]}`

	subquotation := `{"rid":"4","req":"Sub","args":["index_BTC"]}`
	//subkline := `{"rid":"4","req":"Sub","args":["kline_1m_BTC"]}`
	//subPosition := `{"rid":"4","req":"Sub","args":["position_TESTUSD"]}`
	//subPing := `ping`

	//history := `{"req":"GetHistKLine","rid":"103","args":{"Sym":"BTC","Typ":"5m","Offset":0,"Sec":1588416631,"Count":1},"expires":"1588848662118"}`
	//history := `{"req":"GetHistKLine","rid":"103","args":{"Sym":"GMEX_CI_BTC","Typ":"1m","Offset":0,"Sec":1588416631,"Count":1},"expires":"1588848662118"}`
	//subquotation := `{"rid":"88","req":"Sub","args":["index_GMEX_CI_BTC","__slow__"],"expires":"1588848662293"}`
	//unsub := `{"rid":"2","req":"UnSub","args":["kline_1m_BTC"],"expires":"1588848658882"}`

	//writeMessage(c, []byte(history))
	//writeMessage(c, []byte(subkline))
	writeMessage(c, []byte(subquotation))
	//writeMessage(c, []byte(subPosition))

	wg.Wait()
}
