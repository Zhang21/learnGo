package main

import (
	"fmt"
	"time"
)

// We can use channels to synchronize execution across goroutines.

// worker function we'll run in a goroutine.
// The done channel will be used to notify annother goroutine that this functions's work is done.
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	// Start a worker goroutine, giving it the channel to notiry on.
	// Block until we receive a notification from the worker on the channel.
	go worker(done)

	// If you removed the `<-done` line,
	// the program would exit before the worker even started.
	<-done
}
