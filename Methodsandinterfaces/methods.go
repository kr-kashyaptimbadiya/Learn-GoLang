package main

import (
	"fmt"
	"math"
)

//struct type
type Vertex struct {
	X, Y float64
}
//non-struct type
type MyFloat float64

//non-struct type abs method
func (f MyFloat) Abs() float64 {	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//Struct type abs method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Abs is a regular function 
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Pointer receivers method
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	//v.Scale(10)  // The Scale method must have a pointer 
	//receiver to change the Vertex value declared in the main function. 
	//fmt.Println(v.Abs())   //output : 50 bcz after v.Scale(10) our vertex v = {30,40} so, v.abs()= 50
	fmt.Println(v.Abs()) //Abs is a method 
	fmt.Println(Abs(v)) //Abs is a regular function 
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
