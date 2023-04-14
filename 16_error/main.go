package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func foo(a int) (string, error) {
	if a == 42 {
		return "ok", nil
	}
	return "", errors.New("some error")
}

func main() {
	s, err := foo(42)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(s)
}
