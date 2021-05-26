package main

import (
	"fmt"
	"time"
)

// Using goroutines and channel to implement a worker pool.

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This starts up 3 workerss, initially blocked because there are no jobs yets.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 5 jobs and the close that channel to indicate that's all the worker we have.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work.
	// This also encures that the worker goroutines have finished.
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
