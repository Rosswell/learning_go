/**
cli args are a common way to parametrize execution of programs
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args provides access to raw command-line arguments. Note that the first value in this slice is the path to
	// the program, and os.Args[1:] holds the arguments to the program
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// you can get individual args with normal indexing
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
