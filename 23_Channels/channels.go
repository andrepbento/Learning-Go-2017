package main

import "fmt"

/**
  * Channels are the pipes that connect concurrent
  * goroutines.
 */

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}
