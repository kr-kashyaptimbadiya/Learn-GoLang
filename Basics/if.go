package main

import (
	"fmt"
	"math"
)



func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func pow(x, n, lim float64) float64 {
	//If with Short Statement v := math.Pow(x, n);
	if v := math.Pow(x, n); v < lim {
		return v
		//write else immediatly after if.... after closing curly braces of if }.
	}else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10),pow(3, 3, 20))
	
}
