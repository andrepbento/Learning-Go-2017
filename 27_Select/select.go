package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	go func() {
		time.Sleep(time.Second * 5)
		c3 <- "five"
	}()

	for i := 0; i < 3; i++ {
		// Blocking code. Wait's for all three channels
		// to finish their processing
		select {
		case ms1 := <-c1:
			fmt.Println("received", ms1)
		case ms2 := <-c2:
			fmt.Println("received", ms2)
		case ms3 := <-c3:
			fmt.Println("received", ms3)
		}
	}
}
