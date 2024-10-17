package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	test := List[int]{&List[int]{nil, 3}, 2}
	fmt.Println(test.next)
}
