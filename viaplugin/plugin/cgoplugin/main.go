package main

// #include <stdlib.h>
import "C"
import (
	"bytes"
	"encoding/binary"
)

// Random fills the byte slice with random bytes.
func Random(p []byte) (int, error) {
	n := len(p) / 4
	w := bytes.NewBuffer(p[0:0])
	for i := 0; i < n; i++ {
		v := uint32(C.random())

		if m := len(p) % 4; i == n-1 && m > 0 {
			var ww bytes.Buffer
			binary.Write(&ww, binary.BigEndian, v)
			w.Write(ww.Bytes()[:m])
		} else {
			binary.Write(w, binary.BigEndian, v)
		}
	}

	return w.Len(), nil
}

func main() {
	panic("this should be built as go plugin")
}
