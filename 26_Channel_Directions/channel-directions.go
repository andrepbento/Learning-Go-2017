package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pings channel is for sending values
// pongs channel is for receiving values
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
