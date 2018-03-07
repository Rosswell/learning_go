package main

// mutexes can be used to manage more complex states between goroutines
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// state will be a map
	var state = make(map[int]int)

	// this mutex synchronizes access to state, such that operations cannot occur at the same time. The purpose of this
	// is to prevent race conditions
	var mutex = &sync.Mutex{}

	// keeping track of how many read and write operations are performed
	var readOps uint64
	var writeOps uint64

	// starting 100 goroutines to execute repeated reads against the state, once per ms in each goroutine
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				// for each read we pick a key to access, Lock() the mutex to ensure exclusive access to the state,
				// read the value at the chosen key, Unlock() the mutex, and increment the readOps count
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	// pretty much the same looping as the reads, but doing it with writes
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// allow the 10 goroutines to work on the state and mutex for a second
	time.Sleep(time.Second)

	// take and report final ops counts. Executes around 90,000 total operations against the mutex-synchronized state
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
