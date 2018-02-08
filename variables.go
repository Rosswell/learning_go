package main

import "fmt"

func main() {

	// type is inferred
	var a = "initial"
	fmt.Println(a)

	// var allows multiple declarations in the same statement
	var b, c int = 1, 2
	fmt.Println(b, c)

	// type is still inferred
	var d = true
	fmt.Println(d)

	// variables without an initial value are zero-valued, which is 0 for ints
	var e int
	fmt.Println(e)

	// := declares the variable, infers the type, and assigns the value. Equivalent to f string = "short"
	f := "short"
	fmt.Println(f)
}
