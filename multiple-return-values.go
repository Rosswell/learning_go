package main

import "fmt"

// function that takes no parameters, and returns 2 ints
func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// ignoring the first return value
	_, c := vals()
	fmt.Println(c)
}
