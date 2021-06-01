package main

import (
	"fmt"
	"strconv"
)

// Parsing numbers from strings is a basic but common task in many programs.
// The built-in package `strconv` provides the number parsing.

func main() {
	// This 64 tells how many bits of precision to parse.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// The 0 means infer the bas from the string.
	// 64 requires that the result fit in 64bits.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)

	l := strconv.Itoa(451)
	fmt.Print(l, "'s type: ")
	fmt.Printf("%T\n", l)
}
