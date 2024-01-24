package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last.
func mul(x, y int) int {
	return x * y
}

// A function can return any number of results.
func subandswap(x, y int) (int, int, int) {
	return x - y, y, x
}

//Named Return Values
func div(x,y int) (z int){
	z = x / y
	return 
}

func main() {

	fmt.Println(add(22, 8))
	fmt.Println(mul(2, 6))

	a, b, c := subandswap(8, 2)
	fmt.Printf("Subtract Result %d && Swap result %d %d.\n", a, b, c)

	fmt.Println(div(6, 2))
}
