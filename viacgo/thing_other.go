// +build !cgo

package main

import (
	"fmt"
	"math/rand"
)

func random() int {
	fmt.Println("go math/rand#Int")
	return rand.Int()
}
