/**
Defer is used to ensure that a function call is performed later in a program's execution, usually for purposes of
cleanup. defer is often used where e.g. ensure and finally would be used in other languages
*/
package main

import (
	"fmt"
	"os"
)

// creating a file
func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

// writing to the file
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

// closing the file
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func main() {
	// immediately after getting a file object with createFile, we defer the closing of that file with closeFile. This
	// will be executed at the end of the enclosing function(main), after writeFile has finished. output confirms the
	// order of operations
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}
