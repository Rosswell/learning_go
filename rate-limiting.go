package main

import "time"
import "fmt"

// rate limiting controls resource utilization and helps maintain a high quality of service. Go can rate limit
// goroutines, channels, and tickers
func main() {
	// limiting incoming requests by serving requests off the requests channel
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// this channel will receive a value every 200 ms. This is the regulator for rate limiting
	limiter := time.Tick(200 * time.Millisecond)

	// by blocking the receive on the limiter channel before serving each request, we're limited to 1 request every
	// 200ms
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// this scheme is to allow short bursts of requests while preserving the overall rate limit. this is done by
	// buffering the limiter channel. the burstyLimiter channel will allow bursts of up to 3 events
	burstyLimiter := make(chan time.Time, 3)

	// filling up the channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// every 200 ms, adding a new value to burstyLimiter, up to its limit of 3
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// simulating 5 additional requests - the first 3 will benefit from burst capacity
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
