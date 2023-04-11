package main

import "fmt"

type Human struct {
	name    string
	weight  float64
	hobbies []string
}

func (h Human) sayName() {
	fmt.Println("My name is", h.name)
}

type Ninja struct {
	Human
	weapon string
}

func main() {
	h := Human{
		name:    "Alex",
		weight:  75,
		hobbies: []string{"painting"},
	}
	h.sayName()

	n := Ninja{
		Human: Human{
			name:    "Max",
			weight:  78,
			hobbies: []string{"fighting", "skateboarding"},
		},
		weapon: "sword",
	}
	n.sayName()

	fmt.Printf("%#v", n)
}
