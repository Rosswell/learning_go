package main

import "time"
import "fmt"

// timeouts are for programs that require bounds for execution time, maybe because of connecting to external resources
func main() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// res awaits the result, <-time.After awaits a value to be sent after the timeout of 1s
	// since select proceeds with the first case that's ready, the timeout case will be proceeded with if the operation
	// takes more than one second
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
	// if a timeout of more than 3 seconds is allowed, then the receive from c2 will succeed and the result will be
	// printed
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
