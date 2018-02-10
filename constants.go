package main

import "fmt"
import "math"

// declares a constant value
const s string = "constant"

func main() {
	// constants can be used anywhere vars can
	fmt.Println(s)

	const n = 5000000

	// arithmetic is performed with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// a numeric constant has no type until it's explicitly given one, as by an explicit cast
	fmt.Println(int64(d))

	// can also be given a type by using it in a context that requires one, such as a variable assignment
	// or function call (math.Sin()) expects a float64
	fmt.Println(math.Sin(n))
}
