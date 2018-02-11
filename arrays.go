package main

import "fmt"

func main() {

	// declaring an array of length 5 that can only contain ints
	// when initializing an array without values, the array is zero-valued
	var a [5]int
	fmt.Println("emp:", a)

	// assignment by index
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	// declaring an array of ints length 5 with values
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 2D array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
