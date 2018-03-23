/**
how to get the number of seconds, ms, or ns since the Unix epoch
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// use time.Now with Unix or UnixNano to get elapsed time since the Unix epoch
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// to get the ms since epoch, need to manually divide from ns
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// can also convert integer seconds or nanoseconds since the epoch into the corresponding time
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
