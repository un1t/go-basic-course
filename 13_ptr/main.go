package main

import "fmt"

func main() {
	var x int
	var p *int
	fmt.Println(x, p)

	x = 55
	p = &x
	fmt.Println(x, p)

	*p += 3
	fmt.Println(x, p)

	a := "hello"
	b := "world"
	swap(&a, &b)

	fmt.Printf("a=%s, b=%s\n", a, b)
}

func swap(s1, s2 *string) {
	*s1, *s2 = *s2, *s1
}
