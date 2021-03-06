/**
built-in file reading
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// most file writing functions require error checking. this helper streamlines this process
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// slurping an entire file's content into memory
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// more granular control over what is read first requires the file to be opened to obtain an os.File value
	f, err := os.Open("/tmp/dat")
	check(err)

	// read up to 5 bytes to be read from the beginning of the file, but also note how many actually were read
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// can also Seek to a known location in the file and Read from there
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// the io package provides some functions that may be helpful for file reading. for example, read like the ones
	// above can be more robustly implemented with ReadAtLeast
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// there is no built-in rewind, but Seek(0,0) accomplishes this
	_, err = f.Seek(0, 0)
	check(err)

	// the bufio package implements a buffered reader that may be useful both for its efficiency with many small reads
	// and because of the additional reading methods it provides
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// close the file when done with it (usually this would be scheduled immediately after opening with defer
	f.Close()
}
