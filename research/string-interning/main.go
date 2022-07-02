package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

/*
type StringHeader struct {
        Data uintptr
        Len  int
}
*/

// stringptr returns a pointer to the string data.
func stringptr(s string) uintptr {
	return (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
}

func main() {
	s1 := "1234"
	s2 := s1[:2] // "12"
	// s1 and s2 are different strings, but they point to the same string data
	fmt.Println(stringptr(s1) == stringptr(s2)) // true

	// But strings generated at runtime are not interned

	s3 := "12"
	s4 := strconv.Itoa(12)
	fmt.Println(stringptr(s3) == stringptr(s4)) // false
}
