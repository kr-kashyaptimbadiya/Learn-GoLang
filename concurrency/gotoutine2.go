package main

import (
   "fmt"
   "time"
)

func main() {
   go func() {
      for i := 1; i <= 4; i++ {
	     fmt.Println("Goroutine 1:", i)
		  time.Sleep(time.Millisecond * 500)
	  }
   }()

   go func() {
      for i := 1; i <= 4; i++ {
	     fmt.Println("Goroutine 2:", i)
         time.Sleep(time.Millisecond * 500)
      }
   }()

   time.Sleep(time.Second * 2)
}