package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type I implements the interface I,
// but we don't need to explicitly declare that it dose so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}