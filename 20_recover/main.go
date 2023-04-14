package main

import (
	"fmt"
)

func foo() {
	panic("Aaa!")
}

func bar() {
	defer func() {
		r := recover()
		fmt.Println("recovered", r)
	}()

	foo()

	fmt.Println("bar")
}

func main() {
	bar()

	fmt.Println("main")
}
