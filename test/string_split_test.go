package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringsSplit(t *testing.T) {
	config := "127.0.0.1:9092, 127.0.0.1:9093, 127.0.0.1:9094"
	brokers := strings.Split(config, ",")
	fmt.Printf("brokers :=== %v\n", brokers)
	fmt.Printf("broker0: == %s\n", brokers[0])

}
