package main

import (
	"fmt"
	"sync"
	"time"
)

// To wait fo multiple goroutines to finish, we can use a waitgroup.

func worker(id int, wg *sync.WaitGroup) {
	// On return, notify the WaitGroup that we're done.
	defer wg.Done()

	fmt.Printf("worker %d starting\n", id)

	// Sleep to simulate an expensive task.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// used to wait for all the goroutine launched here to finish.
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each.
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block until the wg counter goes back to 0
	// all the workers notified they're done.
	wg.Wait()
}
