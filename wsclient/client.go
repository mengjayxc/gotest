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
	c, resp, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:56789/realtime", nil)
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

	subkline := `{"rid":"4","req":"Sub","args":["quotation:BTC-USDT", "klines_1m:BTC-USDT", "klines_3m:BTC-USDT"]}`
	//history := `{"req":"GetHistKLine","rid":"103","args":{"Sym":"GMEX_CI_BTC","Typ":"1m","Offset":0,"Sec":1588416631,"Count":1},"expires":"1588848662118"}`
	//subquotation := `{"rid":"88","req":"Sub","args":["index_GMEX_CI_BTC","__slow__"],"expires":"1588848662293"}`
	//unsub := `{"rid":"2","req":"UnSub","args":["kline_1m_GMEX_CI_BTC","__slow__"],"expires":"1588848658882"}`

	writeMessage(c, []byte(subkline))

	wg.Wait()
}







