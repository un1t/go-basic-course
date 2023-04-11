package main

import "fmt"

func main() {
	var a []int
	a = append(a, 99)
	a = append(a, 33)
	fmt.Println(a)

	var b []int
	for i := 0; i < 10; i++ {
		b = append(b, i*2)
	}
	fmt.Println(b)

	var c = make([]int, 5)
	fmt.Println(c)

	var d = []int{1, 2, 3, 4, 5}
	e := d[0:5]
	f := d[1:4]
	g := d[1:]
	h := d[:4]
	k := d[:]
	fmt.Println(d, e, f, g, h, k)

	d[1] = 99
	d[2] = -10
	fmt.Println(d, e, f, g, h, k)
}
