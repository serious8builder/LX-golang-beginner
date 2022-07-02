package main

import (
	"fmt"
	"strconv"
)

/*
type error interface {
    Error() string
}
*/

func example(input string) {
	i, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)
}

func main() {
	example("42")
	example("AA")
}
