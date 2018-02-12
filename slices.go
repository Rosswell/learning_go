package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("empty slice:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("len:", len(s))
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// append adds an element and returns the modified list
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// slicing slices is [inclusive, exclusive], like python
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// one-liner declaring and filling a slice
	// note: this is the same as declaring and filling an array, so a little confused by that
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 3 defines the number of inner arrays, but the length of inner slices can be variable
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
