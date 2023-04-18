package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Строки в числа
	a, err := strconv.ParseInt("66", 10, 64)
	fmt.Println(a, err)

	b, err := strconv.Atoi("99")
	fmt.Println(b, err)

	var c int
	_, err = fmt.Sscanf("77", "%d", &c)
	fmt.Println(c)

	d, err := strconv.ParseFloat("33.1", 64)
	fmt.Println(d, err)

	// Числа в строки
	s1 := strconv.Itoa(42)
	fmt.Println(s1)

	s2 := strconv.FormatInt(123, 10)
	fmt.Println(s2)

	s3 := fmt.Sprintf("%d", 77)
	fmt.Println(s3)

	s4 := strconv.FormatFloat(0.1, 'f', 2, 64)
	fmt.Println(s4)

	s5 := fmt.Sprintf("%.2f", 123.45678)
	fmt.Println(s5)

}
