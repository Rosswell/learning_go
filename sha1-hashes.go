/**
SHA1 hashes are frequently used to compute short identities for binary or text blobs. the git revision control system
uses SHA1s extensively to identify versioned files and directories. running the program computes the has and prints it
in a human-readable hex format. you can compute other hashes using a similar pattern to the one shown above. for example
to compute the MD5 hashes import crypto/md5 and use md5.New().
*/
package main

import (
	// several hashing functions are in the crypto package
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	// the patten for generating a hash is sha1.New(), sha1.Write(bytes), then sha1.Sum([]byte{}).
	// starting a new hash here
	h := sha1.New()

	// write expects bytes. use []byte(s) to coerce a string to bytes
	h.Write([]byte(s))

	// this gets the finalized hash result as a byte slice. The argument to Sum can be used to append to an existing
	// byte slice: it usually isn't needed
	bs := h.Sum(nil)

	// SHA1 values are often printed in hex, for example in git commits. the %x format verb converts hash results to a
	// hex string
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
