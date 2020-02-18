package main

import "C"
import "sync"

/*
#import <pthreads.h>

*/

var (
	storeMux sync.Mutex
	store    = map[uintptr]interface{}{}
)

// From discussion:
// https://stackoverflow.com/questions/38730989/cgo-go-pointers-in-go-memory
func storeValue(v interface{}) uintptr {
}
