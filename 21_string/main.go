package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s string = "qы"

	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i, r := range s {
		fmt.Println(i, r, string(r))
	}

	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		fmt.Println(i, runes[i])
	}

	runes[0] = 't'
	s2 := string(runes)
	fmt.Println(s2)
}
