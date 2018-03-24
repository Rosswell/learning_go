/**
allows Go programs to intelligently handle Unix signals. For example, might want a server to gracefully shutdown when
it receives a SIGTERM, or a command line tool to stop processing input if it receives a SIGINT. Here's how to
handle signals in Go with channels. SIGINT is sent to programs in the terminal via ctrl-C
*/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Go signal notification wqorks by sending os.Signal values on a channel. Here a channel is being created to
	// receive those notifications (also need to create a second one to notify the program when to exit)
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// signal.Notify registers the given channel to receive notifications of the specified signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// this goroutine executes a blocking receive for signals. when it gets one it'll print it out and then notify the
	// program that it can finish
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// the program will wait here until it gets the expected signal (as indicated by the goroutiner above sending a
	// value on done) and then exit
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
