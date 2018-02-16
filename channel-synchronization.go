package main

import "fmt"
import "time"

// here is an example of using a blocking receive to wait for a goroutine to finish
// this function uses the done channel to notify another goroutine that it's completed
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	worker(done)

	// putting the done receiver here blocks the main function from executing anything beyond the receiver until it
	// receives a notification from the worker
	<-done
	fmt.Println("main unblocked")
}
