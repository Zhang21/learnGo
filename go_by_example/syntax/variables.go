package main

import "fmt"

// in go, variables are explicitly declared and used by the compiler to e.g. check-correctness of function calls.
// var declares 1 or more vars.
// Go will infer the type of initialized vars.
// Variables declared without a corresponding initialization are zero-valued.
// The := synxtax is shorthand for declaring and initializing a variable.

func main() {
	// implicit declare
	var a = "initial"
	fmt.Println(a)

	// explicit declare
	var b, c int = 1, 2
	fmt.Println(b, c)

	// implicit declare
	var d = true
	fmt.Println(d)

	// uninitial
	var e int
	fmt.Println(e)

	// shorthand
	f := "apple"
	fmt.Println(f)
}
