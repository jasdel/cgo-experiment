package main

/*
#cgo LDFLAGS: -L../bin -lRandom

#include "../libRandom/callback.h"
#include <stdio.h>

extern int my_callback(int id, int i);

static inline int _callback(int id, int i){
	return my_callback(id, i);
}

static inline void* _init_work(int id, int calls) {
	return init_work(id, calls, _callback);
}
*/
import "C"
import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

type slot struct {
	sync.Mutex
	Counter  int
	Worker   unsafe.Pointer
	LastTime time.Time

	TotalDur time.Duration
}

var slots []*slot

//export my_callback
func my_callback(id, i C.int) C.int {
	slot := slots[int(id)]

	//slot.Lock()
	now := time.Now()

	if slot.Counter != 0 {
		slot.TotalDur += now.Sub(slot.LastTime)
	}

	slot.LastTime = now
	slot.Counter++

	//slot.Unlock()

	return i * 2
}

func runCallbacks(atOnce, calls int) {
	slots = make([]*slot, atOnce)

	for i := 0; i < len(slots); i++ {
		slots[i] = &slot{}
		slots[i].Worker = C._init_work(C.int(i), C.int(calls))
	}

	for i := 0; i < len(slots); i++ {
		slot := slots[i]
		C.cleanup(slot.Worker)
		fmt.Println(i, "callback timing from Go point of view", slot.TotalDur/time.Duration(slot.Counter))
	}
}
