/**
parsing numbers from strings
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 64 tells how many bits to parse
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 0 means infer the base from the string. 64 requires that the result fit in 64 bits
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt recognizes hex-formatted numbers
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// ParseUint is also available
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi is a convenience function for basic base-10 int parsing
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parse functions return an error on bad input
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
