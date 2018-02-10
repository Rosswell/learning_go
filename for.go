package main

import "fmt"

func main() {

	// basic for loop, with single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classical intitial/condition/after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for loop without a condition will loop forever until break or return from the enclosing function
	for {
		fmt.Println("loop")
		break
	}

	// continue moves on to the next iteration of the loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
