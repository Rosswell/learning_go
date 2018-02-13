package main

import "fmt"

// a struct is basically a typed collection of fields, that are useful for grouping data together to form records
type person struct {
	name string
	age  int
}

func main() {

	// initializing a new struct
	fmt.Println(person{"Bob", 20})

	// field naming is optional, but preferred
	fmt.Println(person{name: "Alice", age: 30})

	// omitted fields are zero-valued
	fmt.Println(person{name: "Fred"})

	// this prints the pointer of the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// dot notation accesses the struct fields
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// when using dots with struct pointers, the pointers are automatically dereferenced. I think that this basically
	// means that prints the value of the pointer rather than the address.
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
