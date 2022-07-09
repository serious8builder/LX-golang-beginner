package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) AbsPtr() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFuncPtr(v *Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	fmt.Println(v.AbsPtr())     // (&p).AbsPt()
	fmt.Println(AbsFuncPtr(&v)) // must be pointer type

	p := &Vertex{3, 4}
	fmt.Println(p.Abs())     //  (*p).Abs()
	fmt.Println(AbsFunc(*p)) // must be value type?

	fmt.Println(p.AbsPtr())
	fmt.Println(AbsFuncPtr(p))

	fmt.Println(*p)
}
