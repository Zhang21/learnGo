package main

import "fmt"

// This feature is used to in idiomatic go, for example to return both result and error values from a function.

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
