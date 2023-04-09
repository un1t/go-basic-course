package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 55
	x[1] = 3
	x[4] = 7
	fmt.Println(x)

	var y = [3]int{5, 6, 1}
	fmt.Println(y)

	for i := 0; i < 3; i++ {
		fmt.Println(y[i])
	}

	for i := 0; i < len(y); i++ {
		fmt.Println(y[i])
	}

	for _, v := range y {
		fmt.Println(v)
	}

	var m = [3][3]int{
		{3, 4, 1},
		{5, 1},
	}

	for _, row := range m {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
