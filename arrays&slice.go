package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//slice
	//An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
	//flexible view into the elements of an array. In practice, slices are much more common than arrays.
	var s []int = primes[1:4]
	fmt.Println(s)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(st)

	slice := []int{2, 3, 5, 7, 11, 13}

	slice = slice[1:4]
	fmt.Println(slice) //now our original slice changed (3,5,7)

	slice = slice[:2]  // for this operation our slice is(2,3,7)
	fmt.Println(slice) // so, output (3,5)

	slice = slice[1:]  // for this operation our slice is(3,5)
	fmt.Println(slice) // so, output(5)

	//** append Slice
	slice = append(slice, 4, 5, 6)
	fmt.Println(slice)

	slice = slice[:0]
	slice = append(slice, 1, 2, 4, 8, 16, 32)
	//**Range in Slice give  i = index, v = values
	for i, v := range slice {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//Nil slices The zero value of a slice is nil.
	//A nil slice has a length and capacity of 0 and has no underlying array.

	//slice create using make function
	createusingmake := make([]int, 5)
	printSlice("createusingmake", createusingmake)
}
