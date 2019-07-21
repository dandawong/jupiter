package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.Tick(5000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
