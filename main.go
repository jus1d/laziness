package main

import (
	"fmt"
	"lazy/lazy"
)

func main() {
	// Lazy booleans
	fmt.Println("-- Lazy Booleans --")

	t := lazy.Traced(true)
	f := lazy.Traced(false)

	// if operation's result is obvious even with left operand, second one will be never touched
	// (no matter what is the second operand of OR, if first one is true)
	fmt.Printf("  true  || false -> %t\n", lazy.Or(t, f)())

	// same sitaution, if first operand of AND operation is false, the result will never be true
	fmt.Printf("  false || true  -> %t\n", lazy.And(f, t)())

	// opposite situation, with only false, it is not obvious what is the final result of OR operation
	fmt.Printf("  false || true  -> %t\n", lazy.Or(f, t)())
}
