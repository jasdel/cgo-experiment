package main

import "crypto/rand"

// Random fills the byte slice with random bytes.
func Random(p []byte) (int, error) {
	return rand.Read(p)
}

func main() {
	panic("this should be built as go plugin")
}
