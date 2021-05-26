package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// For more complex state we can use a mutex to safely access data across multiple goroutines.
// 在一个goroutine获得Mutex后，其它goroutine只能等到这个goroutine释放该Mutex
func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var readOps uint64
	var writeOps uint64

	// Here we start 100 goroutines to execute repeated reads against the state,
	// once per ms in each goroutines.
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				// For each read we pick a key to access
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Start 10 goroutines to simulate writes.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let the 10 goroutines work on the `state` and `mutex` for a second.
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// With a final lock of `state`, show how it ended up.
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
