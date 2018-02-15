package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// normal call of the f function, running synchronously
	f("direct")

	// creates a new goroutine that executes concurrently with the calling one
	go f("goroutine")

	// creating a new goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Scanln requires a key is pressed before the program exits
	// since the other functions are running asynchronously in separate goroutines, the execution of the main function
	// falls to this point. thus the print statements are interwoven with that of the main function
	fmt.Scanln()
	fmt.Println("done")
}
