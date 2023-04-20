package main

import (
	"fmt"
	"math/rand"
	"time"
)

func writer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		v := rand.Intn(100)
		ch <- v
		time.Sleep(time.Millisecond * 200)
	}
	close(ch)
}

func reader(ch <-chan int) {
	for v := range ch {
		fmt.Println("received", v)
	}
}

func main() {
	ch := make(chan int)
	go writer(ch)
	reader(ch)
}
