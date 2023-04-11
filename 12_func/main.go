package main

import "fmt"

func main() {
	a, b := bar(5)
	fmt.Println(a, b)

	m := moo()
	fmt.Println(m)

	s1 := sum(4, 1, 7)
	fmt.Println(s1)

	var q = []int{4, 1, 2}
	s2 := sum(q...)
	fmt.Println(s2)
}

func foo(x int, s string) int {
	return x + 1
}

func bar(x int) (int, int) {
	return x + 1, x * x
}

func moo() (result int) {
	result = 7
	return
}

func sum(a ...int) int {
	total := 0
	for _, x := range a {
		total += x
	}
	return total
}
