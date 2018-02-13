package main

import "fmt"

// takes an arbitrary number of ints as parameters
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// instead of being forced to feed an array or slice, can just provide ints
	sum(1, 2)
	sum(1, 2, 3)

	// can also use a slice/array as a parameter, using (slice...) notation
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
