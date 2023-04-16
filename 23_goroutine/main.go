package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}
}

func chars() {
	for c := 'a'; c < 'k'; c++ {
		fmt.Println(string(c))
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go chars()
	go numbers()

	time.Sleep(time.Second * 2)
}
