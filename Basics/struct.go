package main

import "fmt"

type Vertex struct {
	X int
	Y int
}
var (
	v1 = Vertex{2, 1}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	poi  = &Vertex{1, 2} // has type *Vertex
)


func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	//Struct fields can be accessed through a struct pointer.
	p := &v
	(*p).X = 1e9 //p.X = 1e9 both are same but (*p).X = 1e9 it looks like a weired.
	fmt.Println(v)

	fmt.Println(v1, poi, *poi, v2, v3)

}
