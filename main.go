package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
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

func main() {
	value := "0x0"

	t1 := "0x5f51ed2a"

	height := "0xa4b24b"

	value2, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}

	timestamp, err := strconv.ParseInt(t1, 0, 64)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}

	height1, err := strconv.ParseInt(height, 0, 64)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}

	fmt.Printf("value2:%v\n", value2)
	fmt.Printf("value = %v\n", value)
	fmt.Printf("timestamp = %v\n", timestamp)
	fmt.Printf("height1 = %v\n", height1)

	add := "19WnXVXAtbGqYKGxrZV7qDWXrudySAcSku"
	aLen := len(add)
	tx := "dbefb28e557a417ed9f0b13dc1973e0445ffa8c50f09bdeb3e509aa3bfbdf1a4"
	txl := len(tx)

	fmt.Printf("alen: +++ %v\n", aLen)
	fmt.Printf("tx: ++ %v\n", txl)

	// 增发
	// MethodID: "0xcbbd8a04"
	h := crypto.Keccak256Hash([]byte(`incrementSupply(uint256)`))
	mintHead := h.Hex()[:10]

	// 销毁
	// MethodID: "0x6d1b229d"
	h = crypto.Keccak256Hash([]byte(`burnTokens(uint256)`))
	burnHead := h.Hex()[:10]

	fmt.Printf("mintHead: +++ %s\n", mintHead)
	fmt.Printf("burnHead: +++ %s\n", burnHead)

}
