package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
)

// 十六进制转换为十进制
func DecodeBigFromHex(hexStr string) (*big.Int, error) {
	hexStr = TrimLeadingZero(hexStr)
	return hexutil.DecodeBig(hexStr)
}

// 删除十六进制字符串前缀多余的0
func TrimLeadingZero(hexStr string) string {
	if strings.Contains(hexStr, "0x") {
		hexStr = strings.TrimPrefix(hexStr, "0x")
	}
	for {
		if !strings.HasPrefix(hexStr, "0") { //不再拥有“0”前缀时跳出
			break
		}
		hexStr = strings.TrimPrefix(hexStr, "0")
	}
	if hexStr == "" {
		return "0x" + "0"
	}
	return "0x" + hexStr
}

func main() {
	input := "0xcbbd8a040000000000000000000000000000000000000000000000000000000001312d00"
	//str := input[0:10]
	//value := input[11:74]
	value := input[11:74]

	value2 := fmt.Sprintf("0x%s", value)

	amount, err := DecodeBigFromHex(value2)
	if err != nil {
		fmt.Printf("decode, err:%s\n", err.Error())
		return
	}

	amount2 := amount.String()

	fmt.Printf("value: %v\n", value)
	fmt.Printf("value2: %v\n", value2)
	fmt.Printf("amount2: %v\n", amount2)
}
