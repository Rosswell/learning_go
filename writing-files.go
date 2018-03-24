/**
writing files; similar patterns as with reading files
*/
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// dumping a string (or just bytes) into a file
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// open a file for writing for more granular writes
	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	// writing byte slices
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// WriteString is also available
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// issue a Sync to flush write to stable storage
	f.Sync()

	// bufio provides buffered writers in addition to the buffered readers we saw earlier
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// use flush to ensure all buffered operations have been applied to the underlying writer
	w.Flush()
}
