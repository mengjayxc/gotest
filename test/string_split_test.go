package test

import (
	"strings"
	"testing"
)

func TestStringsSplit(t *testing.T) {
	//config := "127.0.0.1:9092, 127.0.0.1:9093, 127.0.0.1:9094"
	//newConfig := strings.TrimSpace(config)
	//fmt.Printf("newConfig:%s\n", newConfig)
	//brokers := strings.Split(newConfig, ",")
	//fmt.Printf("brokers :=== %v\n", brokers)
	//fmt.Printf("broker0: == %s\n", brokers[0])
	//fmt.Printf("broker1: == %s\n", brokers[1])
	//fmt.Printf("broker2: == %s\n", brokers[2])
	broker := " 127.0.0.1:9093"
	token := "test_123456789"
	uuid := strings.TrimPrefix(token, "test_")
	t.Logf("uuid:%s", uuid)

	newB := strings.TrimSpace(broker)
	t.Logf("newBroker:%s", newB)
}
