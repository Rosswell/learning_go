package main

import "fmt"

// latter int is declaring the return type
func plus(a int, b int) int {
	// explicit returns are required
	return a + b
}

// grouping consecutively like-typed parameters
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
