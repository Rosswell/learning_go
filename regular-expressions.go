/**
exercising the built-in support for regex functions
*/
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// testing matching of a string
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// for all other regexp tasks, we'll need to Compile an optimized Regexp struct
	r, _ := regexp.Compile("p([a-z]+)ch")

	// same as the match test above
	fmt.Println(r.MatchString("peach"))

	// finds the match
	fmt.Println(r.FindString("peach punch"))

	// finds the first match, but returns the start and end indexes for the match instead of the matching text
	fmt.Println(r.FindStringIndex("peach punch"))

	// includes information about both the whole-pattern matches and teh submatches within those matches. This will
	// return information for both p([a-z]+) and ([a-z]+)
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// returns information about the indexes of matches and submatches - start and end indexes of matches and submatches
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// the All variant of all these tasks provides functions for all matches in the input, rather than just the first.
	// the negative int indicates that all matches should be considered. a positive int indicates that the number of
	// matches should be limited to that number
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// byte slices can also be used, in which case the 'String' part of the regexp function is dropped
	fmt.Println(r.Match([]byte("peach")))

	// the MustCompile function works for creating constants, because Compile returns 2 values
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// regexp can also be used for replacement
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// the func variant allow transformation of matched text with a given function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
