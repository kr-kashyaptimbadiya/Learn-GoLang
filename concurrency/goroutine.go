package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%s %d\n", s, i)
	}
}

func main() {
	go say("world")
	go say("hello")
	for i := 0; i < 5; i++ {
		time.Sleep(5000 * time.Millisecond)
		fmt.Printf("goroutine(). %d\n", i)
	}
}

/*
	NOTE: Here go say is the child thread. So the world may appear at max
	5 times. As it is the child thread, it can be terminated when the main
	thread is terminated.

	Output [Output will keep on varyin coz of context switch]

*/