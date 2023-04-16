package main

import (
	"fmt"
	"sync"
	"time"
)

func write(x *int, mu *sync.RWMutex) {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		*x += 1
		mu.Unlock()
	}
}

func read(x *int, mu *sync.RWMutex) {
	for i := 0; i < 1000; i++ {
		mu.RLock()
		fmt.Println(*x)
		mu.RUnlock()
	}
}

func main() {
	var a = 0
	var mu sync.RWMutex

	go write(&a, &mu)
	go read(&a, &mu)
	go read(&a, &mu)
	go read(&a, &mu)

	time.Sleep(time.Millisecond * 100)
	fmt.Println(a)
}
