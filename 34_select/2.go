package main

import (
	"fmt"
	"time"
)

func worker1(ch chan string) {
	time.Sleep(time.Millisecond * 1000)
	ch <- "done"
}

func worker2(ch chan int) {
	time.Sleep(time.Millisecond * 200)
	ch <- 42
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan int)

	go worker1(ch1)
	go worker2(ch2)

	for i := 0; i < 20; i++ {
		select {
		case val := <-ch1:
			fmt.Println("received", val)
		case val := <-ch2:
			fmt.Println("received", val)
		default:
			fmt.Println(".")
			time.Sleep(time.Millisecond * 100)
		}
	}

}
