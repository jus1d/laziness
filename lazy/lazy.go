package lazy

import (
	"fmt"
	"time"
)

// Lazy wrapper for any default type
type Lazy[T any] func() T

// From constructs a lazy value from default types
func From[T any](value T) Lazy[T] {
	return func() T {
		return value
	}
}

// Traced constructs a lazy value from default types with logging, when the value touched
func Traced[T any](value T) Lazy[T] {
	return func() T {
		fmt.Println("TRACE: touched value: ", value)
		return value
	}
}

// Hang freezed the main thread, so you can clearly see when it's touched
func Hang[T any]() T {
	for {
		fmt.Println("Just hanging around...")
		time.Sleep(500 * time.Millisecond)
	}
}

// First returns first value of pair
func First[T any](a, b Lazy[T]) Lazy[T] {
	return a
}

// Second returns second value of pair
func Second[T any](a, b Lazy[T]) Lazy[T] {
	return b
}

// Or represents logical 'or' for lazy booleans
func Or(a, b Lazy[bool]) Lazy[bool] {
	// yeah, in pure functional style we can't use if-elses, but since in Go there is
	// no ternary operator (even like in Rust) - we'll use if-else
	if a() {
		return From(true)
	}
	return b
}

// And represents logical 'and' for lazy booleans
func And(a, b Lazy[bool]) Lazy[bool] {
	if !a() {
		return From(false)
	}
	return b
}

// Sum represents arithmetical sum for lazy integers
func Sum(a, b Lazy[int]) Lazy[int] {
	return func() int {
		return a() + b()
	}
}
