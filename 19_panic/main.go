package main

func foo() {
	panic("Aaa!")
}

func bar() {
	foo()
}

func main() {
	bar()
}
