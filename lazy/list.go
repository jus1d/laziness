package lazy

import "fmt"

// List is a lazy list basically
type List[T any] Lazy[*InnerList[T]]

// InnterList stands for inner structure of List, to make it lazy
type InnerList[T any] struct {
	Head Lazy[T]
	Tail List[T]
}

// ListFromSlice takes a slice of any type, and convert it to lazy list
func ListFromSlice[T any](xs ...T) List[T] {
	if len(xs) == 0 {
		return func() *InnerList[T] {
			return nil
		}
	}

	list := InnerList[T]{
		Head: From(xs[0]),
		Tail: ListFromSlice(xs[1:]...),
	}

	return func() *InnerList[T] {
		return &list
	}
}

// ListFromRange creates a lazy list from range. Both edges are included
func ListFromRange(start, end Lazy[int]) List[int] {
	val := start()

	if val > end() {
		return func() *InnerList[int] { return nil }
	}

	list := InnerList[int]{
		Head: From(val),
		Tail: ListFromRange(From(val+1), end),
	}

	return func() *InnerList[int] { return &list }
}

// Print just prints the lazy list
func (l List[T]) Print() {
	cur := l()
	for cur != nil {
		fmt.Println(cur.Head())
		cur = cur.Tail()
	}
}
