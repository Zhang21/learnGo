package main

import (
	"fmt"
)

// Closing a channel indicates that no more values will be sent on it.
// This can be useful to communicate completion to the channel's receivers.

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// more will be false if jobs has been closed and all values in the channel have already been received.
			// more ture or false
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		// time.Sleep(time.Second)
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
