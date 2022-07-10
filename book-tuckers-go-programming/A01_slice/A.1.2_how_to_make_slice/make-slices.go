package main

import "fmt"

func main() {
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var s1 []int = arr[1:5]
	var s2 []int = s1[1:8:9]
	var s3 []int = make([]int, 5)
	var s4 []int = make([]int, 0)
	var s5 []int = []int{1, 2, 3, 4, 5}
	var s6 []int

	s4 = append(s4, 10)

	desc(s1)
	desc(s2)
	desc(s3)
	desc(s4)
	desc(s5)
	desc(s6)
}

func desc(s []int) {
	if s == nil {
		fmt.Printf("Len: %v, Cap: %v, Slice: %v\n", len(s), cap(s), "<nil>")
		return
	}
	fmt.Printf("Len: %v, Cap: %v, Slice: %v\n", len(s), cap(s), s)
}
