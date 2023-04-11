package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	foo(a)
	fmt.Println("a=", a)

	var b = []int{1, 2, 3, 4, 5}
	bar(b)
	fmt.Println("b=", b)
}

func foo(x [5]int) {
	x[0] = 99
	fmt.Println("x=", x)
}

func bar(x []int) {
	x[0] = 99
	fmt.Println("x=", x)
}
