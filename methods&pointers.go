package main

import "fmt"
import "math"

type Verte struct {
	X, Y float64
}

func (v *Verte) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Verte, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Verte) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Verte) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Verte{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Verte{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)

	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	q := &Verte{4, 3}
	fmt.Println(q.Abs())
	fmt.Println(AbsFunc(*q))
}




