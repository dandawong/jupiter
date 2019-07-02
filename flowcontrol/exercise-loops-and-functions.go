package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	eps := 0.000001

	abs := func(x float64) float64 {
		if x < 0 {
			return -x
		}
		return x
	}

	for abs(z*z-x) > eps {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
