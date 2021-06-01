package main

import (
	"crypto/sha1"
	"fmt"
)

// SHA1 hashed are frequently used to compute short identities for binary or text blobs.
// For example, the git revision control system.
// Go implements several hash function in various `crypto/*` packages.

func main() {
	s := "sha1 this string"

	// The pattern for generating a hash is `sha1.New()`, `sha1.Write(bytes)`,
	// then `sha1.Sum([]byte{})`.
	h := sha1.New()

	// `Write` expects bytes.
	// If you have a string `s`, use `[]bytes(s)` to create it to bytes.
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte silice.
	bs := h.Sum(nil)

	// SHA1 values are often printed in hex, for example in git commits.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
