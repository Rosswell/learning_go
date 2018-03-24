/**
os.Exit to immediately exit with a given status. unlike C, Go does not use an integer return value from main to indicate
exit status. If you'd like to exit with a non-zero status, os.Exit is necessary. `echo $?` can be used to see the exit
code
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	// defers will not be run when using os.Exit, so this fmt.Println will never be called
	defer fmt.Println("!")

	// exit with status 3
	os.Exit(3)
}
