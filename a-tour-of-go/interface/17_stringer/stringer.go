package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{Name: "A", Age: 65}
	z := Person{Name: "Z", Age: 90}
	fmt.Println(a, z)
}
