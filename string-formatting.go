/**
Go has built-in support for string output formatting with printf
*/
package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var r = fmt.Printf

func main() {

	p := point{1, 2}

	// prints an instance of the point struct
	r("%v\n", p)

	// if the value is a struct, the %+v variant will include the struct's field names
	r("%+v\n", p)

	// the %#v variant prints a Go syntax representation of the value, ie. the source code snippet that would produce
	// that value
	r("%#v\n", p)

	// printing the type of the value
	r("%T\n", p)

	// formatting booleans
	r("%t\n", true)

	// %d offers standard, base-10 formatting for integers
	r("%d\n", 123)

	// binary representation
	r("%b\n", 14)

	// prints the character corresponding to the given integer
	r("%c\n", 33)

	// %x provides hex encoding
	r("%x\n", 456)

	// basic decimal formatting with %f
	r("%f\n", 78.9)

	// %e and %E format the float in slightly different forms of scientific notation
	r("%e\n", 12340000.0)
	r("%E\n", 12340000.0)

	// basic string printing
	r("%s\n", "\"string\"")

	// %q double-quotes strings
	r("%q\n", "\"string\"")

	// %x renders the string in base-16, with two output characters per byte of input
	r("%x\n", "hex this")

	// printing a representation of a pointer
	r("%p\n", &p)

	// controlling the width and precision of the resulting figure. the number after the % in the verb specifies the
	// width. By default, the result will be right-justified and padded with spaces
	r("|%6d|%6d|\n", 12, 345)

	// specifying the width of floats, while also specifying the decimal precision at the same time with the
	// width.precision syntax
	r("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// to left-justify, use the - flag
	r("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// controlling the width when formatting strings, especially to ensure that they align in table-like output.
	// this is for basic left-justified width
	r("|%6s|%6s|\n", "foo", "b")

	// to left-justify, use the - flag, as with numbers
	r("|%-6s|%-6s|\n", "foo", "b")

	// Sprintf formats and returns a string without printing it anywhere
	s := fmt.Sprintf("a %s", "sting")
	fmt.Println(s)

	// formatting and printing to io.Writers, rather than os.Stdout
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
