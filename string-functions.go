/**
general built-in string related functions
*/
package main

import (
	"fmt"
	s "strings"
)

// aliasing prints since it's used a lot in the following funcs
var p = fmt.Println

// since these functions are from the strings package rather than the string object itself, the string needs to be
// passed to the function each time
func main() {
	p("Contains:   ", s.Contains("test", "es"))
	p("Count:      ", s.Count("test", "t"))
	p("HasPrefix:  ", s.HasPrefix("test", "te"))
	p("HasSuffix:  ", s.HasSuffix("test", "st"))
	p("Index:      ", s.Index("test", "e"))
	p("Join:       ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:     ", s.Repeat("a", 5))
	p("Replace:    ", s.Replace("foo", "o", "0", -1))
	p("Replace:    ", s.Replace("foo", "o", "0", 1))
	p("Split:      ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:    ", s.ToLower("TEST"))
	p("ToUpper:    ", s.ToUpper("test"))
	p()

	// while not part of strings, these are the mechanisms for getting the length of a string in bytes and getting a
	// byte by index. len and indexing work at the byte level, which is useful because Go uses UTF-8 encoded strings.
	// working with multi-byte characters will require using encoding-aware operations. See the strings, bytes, runes,
	// and characters packages
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])

}
