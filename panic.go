/**
mostly panics are used to fail fast on errors that shouldn't occur during normal operation, or that we aren't prepared
to handle gracefully. many times, there is one function/package designed to panic specifically. Running this program
will cause it to panic, print an error message and goroutine traces, and exit with a non-zero status. Unlike other
languages, which use exceptions for handling of many errors, in Go it is idiomatic to use error-indicating return values
whenever possible
*/
package main

import "os"

func main() {
	panic("a problem")

	// a common use of panic is to abort if a function returns an error value that we don't know how to (or want to)
	// handle. Here's an example of panicking if we get an unexpected error when creating a new file
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
