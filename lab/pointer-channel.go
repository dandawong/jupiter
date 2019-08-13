package main

import "fmt"

type Vertex struct {
	X, Y float64
}

type IVertex interface {
	Copy() chan IVertex
}

func (src *Vertex) Copy() chan IVertex {
	copy := Vertex(*src)
	var i IVertex = &copy
	c := make(chan IVertex)
	go func() { c <- i }()
	return c
}

func main() {
	var v1 IVertex = &Vertex{1,2}
	v2 := <-v1.Copy()
	fmt.Printf("%p %v\n", &v1, v1)
	fmt.Printf("%p %v\n", &v2, v2)
}
