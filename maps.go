package main

import "fmt"

func main() {

	// general structure: make(map[key-type]val-type)
	m := make(map[string]int)

	// setting values is normal
	m["k1"] = 7
	m["k2"] = 13

	// printing a map shows all its key value pairs
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// len(map) returns the number of key/value pairs
	fmt.Println("len:", len(m))

	// removal of a key/value pair
	delete(m, "k2")
	fmt.Println("map:", m)

	// _ is the blank identifier, which ignores a return value
	// the second, optional return value is a bool that indicates whether a key was present in the map
	// this can be helpful for distinguishing between non-existing keys and keys with zero-values like 0 or ""
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// one-liner declaring and initializing the map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
