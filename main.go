package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GetRandomString(l int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.8f", value), 64)
	return value
}

type BTCAsset struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}

type AssetList struct {
	Action      string    `json:"action"`
	RequestTime time.Time `json:"request_time"`
	Amount      float64   `json:"amount"`
}

type AssetDetail struct {
	BTCChain detail `json:"btc_chain"`
	ETHChain detail `json:"eth_chain"`
}

type detail struct {
	Address string  `json:"address"`
	TxHash  float64 `json:"tx_hash"`
}

func main() {
	//randString :=  GetRandomString(10)
	now := time.Now().UnixNano() / 1e6

	at := now / 100

	at1 := at * 100
	fmt.Printf("now: +++++ %v\n", now)
	fmt.Printf("at: +++++ %v\n", at)
	fmt.Printf("at1: +++++ %v\n", at1)

}
