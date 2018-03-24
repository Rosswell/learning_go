/**
reads input from stdin, processes it, and prints some result to stdout. grep and sed are common line filters
this example writes a capitalized version of all input text. this pattern can be used to write other filters

to test, first make a file with a few lowercase line, the use the line filer to get uppercase lines
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// wrapping the unbuffered os.Stdin with a buffered scanner gives us a convenient Scan method that advances the
	// scanner to the next token, which is the next line in the default scanner
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Text returns the current token, here the next line, from the input
		ucl := strings.ToUpper(scanner.Text())
		// write out the uppercased line
		fmt.Println(ucl)
	}

	// check for errors during Scan. EOF is expected and not reported by Scan as an error
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
