package main

import "fmt"

// channels connect concurrent goroutines. they allow the transfer of information from one goroutine to another
func main() {
	// creating a new channel. channels are typed with the values they convey
	messages := make(chan string)

	// this function sends "ping" to the messages channel, as an async goroutine
	go func() { messages <- "ping" }()

	// the <-channel syntax receives a value from the channel. here it's assigned to the msg variable and printed out
	msg := <-messages
	// by default, the sending and receipt block until both the sender and receiver are ready
	fmt.Println(msg)
}
