package main

import "fmt"

// typical factorial function
func fact(n int) int {
	if n == 0 {
		return 1
	}
	// GoLand interestingly notes the recursive nature and marks the left side to note it
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
