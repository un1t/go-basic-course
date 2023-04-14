package main

import (
	"fmt"
)

func openFile() {
	fmt.Println("open")
}

func writeFile() {
	fmt.Println("write")
}

func closeFile() {
	fmt.Println("close")
}

func main() {
	openFile()
	defer closeFile()

	writeFile()
}
