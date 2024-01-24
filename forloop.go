package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//The init and post statements are optional. 
	// total := 1
	// for ; total < 1000; {
	// 	total += total
	// }
	// fmt.Println(total)

	//bpth ways are right below way looks like a while loop in C
	total := 1
	for  total < 1000 {
		total += total
	}
	fmt.Println(total)


	//If you omit the loop condition it loops forever, so an infinite loop is compactly expressed. 
		//**	for {
		//** 	} 
	//output will be timeout running program
}
