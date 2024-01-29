package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (i IPAddr) String() string{
	temp := ""
	for j := 0; j < 3; j++ {
		temp += strconv.Itoa(int(i[j])) + "."
	}
	temp += strconv.Itoa(int(i[3]))
	return temp
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
