/**
built-in support for times and durations
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// current time
	now := time.Now()
	p(now)

	// building a time struct by providing the year, month, day, etc. they are always associated with a location, ie. UTC
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// extracting the components of the time value as expected
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// extracted as...
	fmt.Printf("%T\n", then.Year())     // int
	fmt.Printf("%T\n", then.Location()) // *time.Location

	// can also do weekdays
	p(then.Weekday())

	// performing time comparisons, resulting in booleans
	p(then.Before(now)) // true
	p(then.After(now))  // false
	p(then.Equal(now))  // false

	// Sub returns a Duration representing the difference between two times. Durations are denoted with h, m, and s
	// within themselves to mark the units
	diff := now.Sub(then)
	p(diff)

	// representing Durations in various units
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// add can advance a time by a given Duration, or subtract from a time with a negative Duration
	p(then.Add(diff))
	p(then.Add(-diff))
}
