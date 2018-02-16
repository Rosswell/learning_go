package main

import "fmt"

// if you're using a channel for a parameter, you can specify whether its to be used for sending or receiving. This
// increases the type-safety of the program

// only accepts a sending channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// accepts one channel for receives (pings), and a second for sends(pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
