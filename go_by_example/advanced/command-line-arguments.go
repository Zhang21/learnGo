package main

import (
	"fmt"
	"os"
)

// Command-line arguments are a common way to parameterize execution of programs.

func main() {
	// `os.Args` provides access to raw cmd-line args.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
