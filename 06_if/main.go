package main

import "fmt"

func main() {
	x := 100

	if x == 50 {
		fmt.Println("foo")
	} else if x == 0 {
		fmt.Println("bar")
	} else {
		fmt.Println("nnn")
	}
}
