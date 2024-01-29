package main

// ** factored import statement
// import (
// 	"fmt"
// 	"math"
// )

//  ** multiple import statements
import "fmt"
import "math"

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
