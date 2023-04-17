package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func increment(x *int64) {
	atomic.AddInt64(x, 1)
}

func main() {
	var a int64 = 0
	for i := 0; i < 1000; i++ {
		go increment(&a)
	}

	time.Sleep(time.Millisecond * 100)
	fmt.Println(a)
}
