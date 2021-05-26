package main

import "fmt"

// By default channels are unbuffered, meaning that they will only accept sends(chan <-) if there is a corresponding receive(<- chan) ready to receive the sent value.
// Buffered channels accept a limited number of values without corresponding receiver for those values.
// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
