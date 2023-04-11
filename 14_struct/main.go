package main

import "fmt"

type Human struct {
	name   string
	weight float64
}

func (h Human) sayName() {
	fmt.Println("My name is", h.name)
}

type Ninja struct {
	Human
	weapon string
}

func (n Ninja) fight() {
	fmt.Println("Fighting...")
}

func main() {
	h := Human{
		name:   "Alex",
		weight: 75,
	}
	h.sayName()

	n := Ninja{
		Human:  Human{name: "Max", weight: 78},
		weapon: "sword",
	}
	n.Human.sayName()
	n.sayName()

	fmt.Printf("%#v\n", n)
	fmt.Println(n.name, n.weight, n.weapon)
}
