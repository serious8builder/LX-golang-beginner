package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	// return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
	s := make([]string, 0, len(ip))
	for _, e := range ip {
		fmt.Println(e)
		s = append(s, strconv.Itoa(int(e)))
	}
	return fmt.Sprintf(strings.Join(s, "."))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":   {127, 0, 0, 1},
		"gogogleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
