package main

import (
	"fmt"
	"sync"
	"time"
)

func increment(c *Container) {
	c.mu.Lock()
	c.a += 1
	c.mu.Unlock()
}

type Container struct {
	a  int
	mu sync.Mutex
}

func main() {
	container := Container{}
	for i := 0; i < 1000; i++ {
		go increment(&container)
	}

	time.Sleep(time.Millisecond * 100)
	fmt.Println(container.a)
}
