package main

import "time"
import "fmt"

// timers are for executing code at some point in the future, or repeatedly at some interval
func main() {

	// timers represent a single event in the future. timers are told how long to wait, and it provides a channel that
	// will be notified at that time. this one waits 2 seconds
	timer1 := time.NewTimer(2 * time.Second)

	// this blocks this timer's channel C until it sends a value indicating that the timer expired
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// one difference between a timer and sleeping is that you can cancel a timer before it expires
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	// only stops if the timer has not had a chance to expire
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}