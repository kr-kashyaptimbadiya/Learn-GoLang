package main

import "fmt"

var c, python, java bool

// A var declaration can include initializers, one per variable.
// initializer is present, the type can be omitted
var cpp, golang = true, "no" //variable take type of initializers cpp - boolean and goland - string

//Outside a function, every statement begins with a keyword (var, func, and so on)
// and so the := construct is not available.
//**z := 6

// ** Constants
const Pi = 3.14

func main() {
	var i int

	var x = 3

	//Inside a function, the := short assignment statement can be used in place of a var
	y := 4

	fmt.Println(i, c, python, java) //output 0, false, false ,fasle, ""
	fmt.Println(x, cpp, golang)
	fmt.Println(y)

	// Variables declared without an explicit initial value are given their zero value.
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	//** Constants
	const World = "Kashyap"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
