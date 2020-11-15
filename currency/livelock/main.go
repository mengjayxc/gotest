package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func cadenBroad(cadence *sync.Cond) {
	for range time.Tick(1 * time.Millisecond) {
		cadence.Broadcast()
	}
}

func takeStep(cadence *sync.Cond) {
	cadence.L.Lock()
	cadence.Wait()
	cadence.L.Unlock()
}

func tryDir(dirnName string, dir *int32, out *bytes.Buffer, takeStep func(cond *sync.Cond), cadence *sync.Cond) bool {
	fmt.Printf("out: %v, %s-111\n", out, dirnName)
	atomic.AddInt32(dir, 1)
	takeStep(cadence)
	if atomic.LoadInt32(dir) == 1 {
		fmt.Printf("out: %v.Success-222\n", out)
		return true
	}
	takeStep(cadence)
	atomic.AddInt32(dir, -1)
	return false
}

// 活锁问题
func main() {
	cadence := sync.NewCond(&sync.Mutex{})
	var left, right int32
	go cadenBroad(cadence)
	tryLeft := func(out *bytes.Buffer) bool {
		return tryDir("left", &left, out, takeStep, cadence)
	}

	tryRight := func(out *bytes.Buffer) bool {
		return tryDir("right", &right, out, takeStep, cadence)
	}

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot-333:\n", name)
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!-444", name)
	}

	var peopleInHallWay sync.WaitGroup
	peopleInHallWay.Add(2)

	go walk(&peopleInHallWay, "Alice")
	go walk(&peopleInHallWay, "Barbara")

	peopleInHallWay.Wait()
}
