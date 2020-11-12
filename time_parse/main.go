package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func Wrap(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

func main() {
	var timestamp int64
	timestamp = 1601365737800
	// 1601365235700
	// 1604295154.153779
	// 2020-09-29 15:40:35.7

	// 1601365966557
	// 2020-09-29 15:52:46.557
	time1 := time.Unix(0, timestamp*int64(time.Millisecond))
	fmt.Printf("time1: ++++++ %v\n", time1)

	id := time.Now().UnixNano() / 1e6
	ids := strconv.FormatInt(id, 10)

	fmt.Printf("id: ++++%v\n", id)
	fmt.Printf("ids: ++++%s\n", ids)

	//将float64转成精确的int64
	ts := 1604295154.153779
	wts := Wrap(ts, 6)
	fmt.Printf("wts: %d\n", wts)

	ms := wts / 1e3 //毫秒

	ts2 := time.Unix(0, wts*int64(time.Microsecond))
	ts3 := time.Unix(0, ms*int64(time.Millisecond))
	fmt.Printf("ts2: %v\n", ts2)
	fmt.Printf("ms: %v\n", ms)
	fmt.Printf("ts3: %v\n", ts3)

}
