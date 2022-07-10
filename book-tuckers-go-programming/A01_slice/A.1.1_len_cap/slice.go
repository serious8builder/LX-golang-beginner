package main

import "fmt"

func main() {
	var array [10]int
	var slice []int = array[1:3]
	fmt.Printf("Len:%v, Cap:%v\n", len(slice), cap(slice)) // 2, 9

	slice = array[1:3:6]
	fmt.Printf("Len:%v, Cap:%v\n", len(slice), cap(slice)) // 2, 5

}
