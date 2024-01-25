package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type MYFLOAT float64

type Vert struct {
	X, Y float64
}

func (f MYFLOAT) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
} 

func (v *Vert) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MYFLOAT(-math.Sqrt2)
	v := Vert{3, 4}

	a = f  // a MYFLOAT implements Abser
	a = &v // a *Vert implements Abser

	// In the following line, v is a Vert (not *Vert)
	// and does NOT implement Abser.
	//a = v

	fmt.Println(a.Abs())
}


