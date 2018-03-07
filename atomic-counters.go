package main

// atomic counters are another mechanism for managing state between goroutines, in addition to communication over channels
import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	// ops represents an always-positive counter
	var ops uint64

	// this loop starts 50 goroutines, each of which increments the counter ~1 ms
	for i := 0; i < 50; i++ {
		go func() {
			for {

				// this atomically increments the ops counter by giving the memory address of it with the & syntax
				atomic.AddUint64(&ops, 1)

				// wait 1ms between increments
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// wait a second to allow ops to be incremented
	time.Sleep(time.Second)

	// this allows the extraction of the ops counter while its actively being incremented by making a copy of the current
	// value. As above, this requires fetching the memory address of ops
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
