package main

import "time"
import "fmt"

// this is the one worker, that is to be used to run several concurrent instances of. the workers receive work through
// the jobs channel, and send the corresponding results on results. The 1 second timer simulates an expensive task
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// in order to use our pool of wokrers we need to send them work and collect their results. We make 2 channels for
	// these purposes
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// this starts up 3 workers, initially blocked because there are no jobs yet
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// here we send 5 jobs and then close that channel to indicate that's all the work we have
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// collecting all the results of the work
	for a := 1; a <= 5; a++ {
		<-results
	}
}
