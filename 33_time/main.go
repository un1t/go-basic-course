package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()

	fmt.Println(t1.Day(), t1.Month(), t1.Year())
	fmt.Println(t1.Format(time.RFC3339))

	t2 := t1.Add(time.Hour * 5)
	t3 := t1.AddDate(0, 0, -3)

	t1.Before(t2)
	t1.After(t3)
	t1.Equal(t2)

	diff := t2.Sub(t1)

	fmt.Println(diff)
	fmt.Println(diff.Hours(), diff.Minutes(), diff.Seconds())
}
