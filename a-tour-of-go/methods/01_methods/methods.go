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

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(n int) {
	v.X *= float64(n)
	v.Y *= float64(n)
}

// Scale 위 함수와 같은 내용
func Scale(v *Vertex, n int) {
	v.X *= float64(n)
	v.Y *= float64(n)
}

// Scale2 value receiver 의 경우 리턴을 안 하면 의미가 없다.
func (v Vertex) Scale2(n int) {
	v.X *= float64(n)
	v.Y *= float64(n)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	v.Scale(2)
	fmt.Println(v) // {6, 8}

	v.Scale2(3)
	fmt.Println(v) // {6, 8}
}
