package main

import "fmt"

func main() {
	var s []int
	a := [0]int{}
	var p = &a
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
	fmt.Println(p, len(p), cap(p))
	fmt.Printf("%T %v\n", p, p)
	fmt.Println(a, len(a), cap(a))
}
