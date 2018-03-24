/**
built-in Go support for base64 encoding/decoding. the output string encodes to slightly different values with the
standard and URL base64 encoders (trailing + vs -), but they both decode to the original string as desired
*/
package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	// string for encoding usage
	data := "abc123!?$*&()'-=@~"

	// go supports both standard and url-compatible base64. here's how to encode using the standard encoder. the encoder
	// requires a []byte so we cast our string to that type
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// decoding may return an error, which you can check if you don't already know the input to be well-formed
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// this encodes/decodes using a url-compatible base64 format
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
