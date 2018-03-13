package main

import (
	"fmt"
	"sort"
)

// go has a builtin sort package. all sorting done this way is done in-place
func main() {
	// sorting strings
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// sorting ints
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:    ", ints)

	// returns a boolean to see if a slice is already sorted
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
