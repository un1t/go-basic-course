package main

import "fmt"

func main() {
	var a int32 = 2023
	var b uint16 = uint16(a)
	var c = float32(b)
	var d = int8(a)
	var s = string(a)

	fmt.Printf("a=%v, b=%v, c=%v, d=%v, s=%v\n", a, b, c, d, s)

}
