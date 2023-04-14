package main

import "fmt"

func main() {
	var i interface{} = 55

	v, ok := i.(int)
	fmt.Printf("%T %v %v\n", v, v, ok)

	switch i.(type) {
	case int, float64:
		fmt.Println("number")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("other")
	}
}
