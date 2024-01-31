package main

import "fmt"

type Ver struct {
	Lat, Long float64
}

// var m map[string]Ver

// func main() {
// 	m = make(map[string]Ver)
// 	m["Bell Labs"] = Ver{
// 		40.68433, -74.39967,
// 	}
// 	fmt.Println(m["Bell Labs"])
// }

//	var m = map[string]Ver{
//		"knackroot": Ver{
//			41.11212, -72.312112,
//		},
//		"Google": Ver{
//			37.42202, -122.08408,
//		},
//	}
var m = map[string]Ver{
	"knackroot": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m["knackroot"])
	fmt.Println(m["Google"])

	mp := make(map[string]int)

	mp["Answer"] = 42
	fmt.Println("The value:", mp["Answer"])

	mp["Answer"] = 48
	fmt.Println("The value:", mp["Answer"])

	v, ok := mp["Answer"] //v = 48 mp["Answer"], ok = true (means Answer Present in Map or not)
	fmt.Println("The value:", v, "Present?", ok)
}
