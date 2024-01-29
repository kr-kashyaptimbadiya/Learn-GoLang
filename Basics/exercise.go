package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	i := 0
	for ; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	fmt.Printf("For %dth iteration, approximated value is %g\n", i+1, z)
	return z
}

func main() {

	for i := 1; i < 2001; i++ {
		fmt.Printf("Sqrt(%d) = %g\n", i, Sqrt(float64(i)))
	}
}
