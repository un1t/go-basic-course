package main

import (
	"log"
	"time"
)

func main() {

	val, err := time.Parse("sdfds", "sdfdsaf")
	if err != nil {
		log.Println("some error")
	}

}
