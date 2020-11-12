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

	c, resp, err := websocket.DefaultDialer.Dial("wss://ss.abkjl.com/v1/market", nil)
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

	// 订阅盖亚 index
	subQuotation := `{"req":"Sub","rid":"11","args":["index_GMEX_CI_ETH"]}`

	// 盖亚的订阅
	// rid:10 klines_1m
	//subkline := `{"req":"Sub","rid":"10","args":["kline_1w_BTC.USDT"],"expires":1593766170283}`

	//writeMessage(c, []byte(subkline))
	writeMessage(c, []byte(subQuotation))

	wg.Wait()
}
