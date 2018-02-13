package main

import "fmt"

// this function returns another function, effectively closes over the variable i to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// creates a function of nextInt
	nextInt := intSeq()

	// each call to nextInt calls the internal function within intSeq(), which iterates and returns i
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// this confirms that the current state of i is unique to nextInt
	newInts := intSeq()
	fmt.Println(newInts())
}
