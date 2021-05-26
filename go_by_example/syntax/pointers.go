package main

import "fmt"

// Go supports pointers, allowing you to pass references to values and records within your program.

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

// zeroval doesn't change the i in main, but zeroprt does
// because it has a reference to the memory address for that variable.
func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i memory address of i
	zeroptr(&i)
	fmt.Println("zeroprt:", i)

	fmt.Println("pointer:", &i)
}
