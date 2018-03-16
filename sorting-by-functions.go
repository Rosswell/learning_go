/*
custom sorting in go, ie. by length or something like that. Using this same pattern of creating a custom type,
implementing the three Interface methods on that type, and then calling sort.Sort on a collection of that custom type,
we can sort Go slices by arbitrary functions
*/
package main

import (
	"fmt"
	"sort"
)

// byLength is just an alias for []string. This will be the corresponding type for the sorting functions
type byLength []string

// implementation of Len, Less, and Swap so sort.Sort can be used. Len is similar across types
func (s byLength) Len() int {
	return len(s)
}

// Swap is also similar across types
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less holds the actual custom sorting logic. In this case, it's sorting in order of increasing string length
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// implementing the custom sort by casting the original fruits slice to byLength, and then using sort.Sort on that typed
// slice
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
