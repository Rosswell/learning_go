/**
specifying options through flags
*/
package main

import (
	"flag"
	"fmt"
)

func main() {
	// basic flag declarations are available for string, integer, and boolean options. here a string flag word is being
	// declared with a default value "foo", and a short desciption. This flag.String function returns a string pointer
	// (not a string value); using the pointer is below
	wordPtr := flag.String("word", "foo", "a string")

	// declaring numb and fork flags, using a similar approach to the word flag
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// it's also possible to declare an option that uses an existing var declared elsewhere in the program. note a
	// pointer needs to be passed to the function
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// once all flags are declared, call flag.Parse() to execute the command-line parsing
	flag.Parse()

	// dumping out the parse options and any trailing positional arguments. note that the pointers need to be
	// dereferenced with eg. *wordPtr to get the actual option values
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	// the flag package requires that all flags appear before positional arguments (otherwise the flags will be
	// interpreted as positional arguments)
	fmt.Println("tail:", flag.Args())
}
