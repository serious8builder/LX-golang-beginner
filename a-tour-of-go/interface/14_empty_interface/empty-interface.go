package main

import (
	"fmt"
	"github.com/buildgo/gohi"
)

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	fmt.Println(gohi.SayHi())
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
