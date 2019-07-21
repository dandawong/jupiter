package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walkHelper func(t *tree.Tree, ch chan int)
	walkHelper = func(t *tree.Tree, ch chan int) {
		if t == nil {
			return
		}
		walkHelper(t.Left, ch)
		ch <-t.Value
		walkHelper(t.Right, ch)
	}

	walkHelper(t, ch)
	defer close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 == false && ok2 == false {
			return true
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
