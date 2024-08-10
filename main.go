package main

import (
	"fmt"
	"lazy/lazy"
)

func main() {
	value := lazy.From(4)

	result := lazy.Second(value, lazy.Hang[int])()
	fmt.Printf("Result: %d\n", result)
}
