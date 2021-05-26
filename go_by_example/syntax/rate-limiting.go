package main

import (
	"fmt"
	"time"
)

// Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service.
// Go elegantly supports rate limiting with goroutines, channels, and tickers.

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This `limiter` channel will receive a value every 200 ms.
	// This is the regulartor in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on a receive from the `limiter` channel before serving each request,
	// we limit 1 request every 200ms.
	for req := range requests {
		fmt.Println("---before limiter blocking---")
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// We may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit.
	// We can accomplish this by buffering our limiter channel.
	// This `burstyLimiter` channel will allow bursts of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200ms we'll try to add a new value to `burstyLimiter`,
	// up to its limits of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		println("###before burstyLimiter blocking###")
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
