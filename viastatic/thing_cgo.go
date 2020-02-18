// +build cgo

package main

// #include <stdlib.h>
import "C"
import "fmt"

func random() int {
	fmt.Println("native C.random")
	return int(C.random())
}
