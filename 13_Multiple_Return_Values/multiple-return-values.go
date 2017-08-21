package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func lotsOfVals() (int, float32, float64, string) {
	return 3, 2.71, 3.14, "fim"
}

func main() {
	a, b := vals()
	fmt.Println("Function vals() returned", a, "and", b)

	_, c := vals()
	fmt.Println("Second returned value of vals():", c)

	d, e, f, g := lotsOfVals()
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
