package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	// FIFO queue order
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// When a channel doesn't have more values in the
	// queue, a fatal error occurs (deadlock)
	//fmt.Println(<-messages)
}
