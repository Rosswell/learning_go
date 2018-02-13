package main

import "fmt"

type rect struct {
	width, height int
}

// area method has a receiver type of the rect pointer. Pointer receiver types avoid copying on method calls, or allow
// methods to mutate the receiving struct
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// values and points are automatically converted for method calls
	rp := &r
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", rp.perim())
}
