package main

import "fmt"

// closing channels is used to indicate that there will be no more messages sent on it. This can be used to communicate
// completion to a channel's receivers
func main() {
	// the jobs channel will be used to communicate work done from the main() goroutine to a worker goroutine. when we
	// have no more jobs for the owkrer we'll close the jobs channel
	jobs := make(chan int, 5)
	done := make(chan bool)

	// this is the worker goroutine. it repeatedly receives from jobs with j, more := <-jobs. In this special 2-value
	// form of receive, the more value will be false if jobs has been closed and all values in the channel have already
	// been received. We use this to notify on done when we've worked all our jobs
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// this sends 3 jobs to the worker over the jobs channel, then closes it
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// we await the worker using the synchronization approach we saw earlier - blocking the main function from
	// completing until this receiver receives a notification from the worker
	<-done
}
