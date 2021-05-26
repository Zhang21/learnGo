package main

import (
	"fmt"
	"time"
)

// Execute go code at some point, or repeatedly at some interval. Go's built-in timer and ticker features make both of these tasks easy.
// Timers represent a single event in the future. You tell the timer how long you want to wait, and it provides a channel that will be notified at that time.
// If you just wanted to wait, you could have used time.Sleep. One reason a timer may be useful is that you can cancel the timer before it fires.

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	fmt.Println("--------")

	<-timer1.C
	fmt.Println("Timer 1 fired")

	// timer2 should be stopped before it has a chance to fire.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Time 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
