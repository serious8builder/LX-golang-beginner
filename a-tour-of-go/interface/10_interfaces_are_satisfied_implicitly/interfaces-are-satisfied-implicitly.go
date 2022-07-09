package main

import "fmt"

type DuckLike interface {
	Quack()
}

type Duck struct {
	Sound string
}

type RubberDuck struct {
}

func (duck *Duck) Quack() {
	fmt.Println(duck.Sound)
}

func (duck RubberDuck) Quack() {
	fmt.Println("...")
}

func main() {
	var duckLike DuckLike = &Duck{"Quack"}
	duckLike.Quack()

	duckLike = RubberDuck{}
	duckLike.Quack()
}
