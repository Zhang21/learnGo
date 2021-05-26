package main

import "fmt"

// Go has various value types including:
// strings, integers, floats, booleans...

func main() {
	// strings
	fmt.Println("go" + "lang")
	// int, float
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	// bool
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
