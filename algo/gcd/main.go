package main

import "fmt"

func main() {
	m, n := 2166, 6099
	fmt.Println("GCD:", m, n, GCD(m, n), gcd(m, n))
}
