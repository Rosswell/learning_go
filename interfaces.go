package main

import "fmt"
import "math"

// interfaces are named collections of method signatures. in order to implement an interfacce in Go, we just need to
// implement all the methods in the interface, in this case area() and perim(). As far as I can tell, it's pretty much
// a class whose methods are defined on a per-struct basis
type geometry interface {
	area() float64
	perim() float64
}

// defining the structs to implement the interface on
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// interface implementation for rects
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// interface implementation for circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// using the interface as the variable type means that you can then call the methods in the interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// since circle and rect both implement the geometry interface, they are both valid inputs to measure()
	measure(r)
	measure(c)
}
