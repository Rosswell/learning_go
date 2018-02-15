package main

import "fmt"

// by default, channels are unbuffered, which means that they will only accept sends (chan <-) if there is a
// corresponding receive (<- chan) ready to receive the sent value. buffered channels accept a limited number of values
// without a corresponding receiver for those values
func main() {
	// this channel can hold 2 values
	messages := make(chan string, 2)

	// filling the channel without a concurrent receive
	messages <- "buffered"
	messages <- "channel"

	// receiving the two messages as normal
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
