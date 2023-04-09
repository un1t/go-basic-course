package main

import "fmt"

func main() {
	var m = make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "a")

	fmt.Println(m)

	c, ok := m["c"]

	fmt.Println(c, ok)

	var x = map[string]int{
		"a": 1,
		"c": 3,
	}

	fmt.Println(x)

	var y = map[string]int{}
	y["q"] = 7

	fmt.Println(y)
}
