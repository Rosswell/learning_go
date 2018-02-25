package main

import "time"
import "fmt"

// if timers are for when you want to do something once in the future, tickers are for when you want to do something
// repeatedly at regular intervals
func main() {

	// iterates over the values as they arrive evey 500ms
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	// tickers can be stopped like timers. once a ticker is stopped, it won't receive any more values on its channel.
	// here, it's stopped after 1600ms, so the ticker should tick 3 times before stopping
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
