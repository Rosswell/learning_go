package main

import "fmt"

func main() {
	// ranges are basically just for iterating over a variety of objects

	nums := []int{2, 3, 4}
	sum := 0
	// summing the numbers in a slice (or array)
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// the first return value is the index, but it can be ignored with _, as above
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// key/value pairs can be iterated over, but the value value is optional
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on a string iterates over Unicode code points, starting at the byte index of the rune and then the rune
	// itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
