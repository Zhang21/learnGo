package main

import "os"

// A `panic` typically means something went unexpectedly wrong.

func main() {
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
