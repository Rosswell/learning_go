package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// implementation of synchronizing state of goroutines using channels rather than mutexes
// each piece of data is owned by exactly one goroutine

// state is owned by a single goroutine, which means that it's data won't be corrupted by other processes
// attempting concurrent access. in order to read or write that state, other goroutines will send messages
// to the owning goroutine and receive corresponding replies. these readOp and writeOp structs encapsulate
// those requests and a way for the owning goroutine to respond
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	var readOps uint64
	var writeOps uint64
	// these read/write channels are to be used by other goroutines to issues read and write requests
	reads := make(chan *readOp)
	writes := make(chan *writeOp)
	// this is the goroutine that owns the state, which is a map that is private to the stateful goroutine
	// this goroutine repeatedly selects on the reads and writes channels, responding to requests as they arrive
	// a response is executed by the first performing the requested operation and then sending a value on the
	// response channel resp to indicate success (and the desired value in the case of reads)
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()
	// this starts 100 goroutines to issue reads to the state-owning goroutine via the reads channel. Each read
	// requires constructing a readOp, sending it over the reads channel, and the receiving the result over the
	// provided resp channel
	for r := 0; r < 100; r++ {
		go func() {
			read := &readOp{
				key:  rand.Intn(5),
				resp: make(chan int)}
			reads <- read
			<-read.resp
			atomic.AddUint64(&readOps, 1)
			time.Sleep(time.Millisecond)
		}()
	}
	// starting 10 writes as well, using a similar approach
	for w := 0; w < 10; w++ {
		go func() {
			write := &writeOp{
				key:  rand.Intn(5),
				val:  rand.Intn(100),
				resp: make(chan bool)}
			writes <- write
			<-write.resp
			atomic.AddUint64(&writeOps, 1)
			time.Sleep(time.Millisecond)
		}()
	}
	// letting the goroutines work for a second. Finally, capturing and reporting the op counts
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
