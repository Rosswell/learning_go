package main

import "fmt"

// this function has a copy of ival distinct from the one in the calling function
func zeroval(ival int) {
	ival = 0
}

// *int means that the function takes an int pointer. The code in the function then dereferences the point from its
// memory address to the current values at that address. Assigning a value to a dereferenced pointer changes the value
// at the referenced address
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i) // 1

	zeroval(i)
	fmt.Println("zeroval:", i) // 1 because i is only 1 within the function, doesn't reach global scope

	// &i is the pointer to i. So basically what's happening is the value at that memory address is being changed,
	// rather than assigning the variable to a different memory address
	zeroptr(&i)
	fmt.Println("zeroptr:", i) // 0; the pointer reaches global scope because its pointer is revalued

	fmt.Println("pointer address:", &i)
}
