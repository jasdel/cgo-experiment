// +build !cgo

package main

import "crypto/rand"

// Random populates the passed in slice with random bytes.
func Random(p []byte) (int, error) {
	return rand.Read(p)
}
