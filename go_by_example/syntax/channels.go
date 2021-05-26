package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines.
// You can send values from one goroutines and receive those values into another goroutine.
// Send a value into a channel use the `channel <-` syntax.
// By default sends and receives block until both the sender and receiver are ready.
// This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}
