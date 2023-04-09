package main

import "fmt"

func main() {
	x := 0
	for x < 10 {
		x++
		if x == 3 {
			continue
		}
		fmt.Println(x)

		if x > 5 {
			break
		}
	}

out:
	for i := 0; i < 5; i += 2 {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break out
			}
			fmt.Println(i, j)
		}
	}
}
