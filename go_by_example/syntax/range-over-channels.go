package main

import (
	"fmt"
)

// We can also use for and range to iterate over values received from a channel.

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
