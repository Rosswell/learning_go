/**
setting, getting, and listing environment variables
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// setting/getting a key/value pair. this will return an empty string if the key isn't present in the environment
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// os.Environ lists all key/value pairs in the environment. this returns a slice of strings in the form KEY=value.
	// the key and value can be retrieved by using strings.Split
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
