/**
parsing urls in Go
*/
package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// example url, which includes a scheme, authentication info, host, port, path, query params, and query fragment
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// parse the url to ensure there are no errors
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// accessing the scheme
	fmt.Println(u.Scheme)

	// User contains all authentication info; call Useername and Password on this for individual values
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// The Host contains both the hsotname and the port, if present. Use SplitHostPort to extract them
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// extracting the fragment after the #
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// to get query params in a string of k=v format, use RawQuery. You can also parse query params into a map. The
	// parse query param maps are from strings to slices of strings, so index into [0] if you only want the first value
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
