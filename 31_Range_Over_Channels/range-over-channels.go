package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// If the channel was not close, the program
	// would fail with a deadlock error
	for elem := range queue {
		fmt.Println(elem)
	}
}
