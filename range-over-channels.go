package main

import "fmt"

// for and range can be using to both iterate over data structures and values received from a channel
func main() {

	// iterating over two values in the queue channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// possible to close a non-empty channel but still have the remaining values be received
	close(queue)

	// iterates over each element as it's received from queue. Because queue was closed above, the iteration stops after
	// receiving the two elements

	for elem := range queue {
		fmt.Println(elem)
	}
}
