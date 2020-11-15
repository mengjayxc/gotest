package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func printSum(v1, v2 *value, wg *sync.WaitGroup) {
	defer wg.Done()
	v1.mu.Lock()
	defer v1.mu.Unlock()

	time.Sleep(2 * time.Second)
	v2.mu.Lock()
	defer v2.mu.Unlock()

	fmt.Printf("sum=%v\n", v1.value+v2.value)
}

// 死锁问题！！！
func main() {
	var wg sync.WaitGroup
	var a = value{
		mu:    sync.Mutex{},
		value: 1,
	}

	var b = value{
		mu:    sync.Mutex{},
		value: 2,
	}

	wg.Add(2)
	go printSum(&a, &b, &wg)
	go printSum(&b, &a, &wg)
	wg.Wait()
}
