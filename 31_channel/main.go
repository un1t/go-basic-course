package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "foo"
		ch <- "bar"
		close(ch)
	}()

	for x := range ch {
		fmt.Println(x)
	}

}
