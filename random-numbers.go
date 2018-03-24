/**
math/rand provides pseudorandom number generation
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// returns a  random int n, 0 <= n < 100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// rand.Float64 returns a float64 f, 0.0 <= f < 1.0
	fmt.Println(rand.Float64())

	// this can be ysed to generate random floats in other ranges, for example 5.0 <= f' < 10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// the default number generator is deteministic, so it'll produce the same sequence of numbers each time by default.
	// to produce varying sequenes, give it a seed that changes. Note that this is not safe to use for random numbers
	// you intend to be secret, use crypto/rand for those
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// the resulting rand.Rand can be used like the functions in the rand package
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// if you seed a source with the same number, it profuces the same sequence of random numbers
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
