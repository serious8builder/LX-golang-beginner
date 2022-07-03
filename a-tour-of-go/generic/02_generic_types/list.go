package main

import (
	"fmt"
)

type List[T any] struct {
	val  T
	next *List[T]
}

func (list *List[T]) Len() int {
	count := 0
	for node := list; node != nil; node = node.next {
		count++
	}
	return count
}

func (list List[T]) Head() T {
	return list.val
}

func (list *List[T]) Append(x T) {
	cur := list
	for ; cur.next != nil; cur = cur.next {
	}
	cur.next = &List[T]{val: x}
}

func (list *List[T]) String() string {
	if list == nil {
		return "nil"
	}

	return fmt.Sprintf("%v->%v", list.val, list.next.String())
}

func main() {
	myList := &List[int]{val: 3}
	myList.Append(5)
	fmt.Printf("Head: %v\n", myList.Head())
	fmt.Printf("Len: %v\n", myList.Len())
	fmt.Println(myList)
}
