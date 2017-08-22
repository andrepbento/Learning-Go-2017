package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

/**
	Goroutines are lightweight thread's of execution.
 */

func main() {
	f("direct")

	go f("go routine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("You have entered:", input)
	fmt.Println("done")
}
