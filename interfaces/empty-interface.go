package main

import "fmt"

type I interface {
}

func main() {
	var i I
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	describe("world")
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
