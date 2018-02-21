package main

import "fmt"

// using select with a default clause implements non-blocking sends, receives, and non-blocking multi-way selects
func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// non-blocking receive - proceeds immediately on either case, based on if messages has a value or not
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// non-blocking send - msg cannot be sent to the messages channel, because the channel has no buffer and there is
	// no receiver. Therefore the default case is selected
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// multi-way-non-blocking select - using multiple cases above the default clause to attempt non-blocking receives
	// on both messages and signals
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
