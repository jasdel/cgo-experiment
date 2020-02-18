package main

import (
	"log"
	"syscall"

	"golang.org/x/sys/windows"
)

// ported from stdlib:
// src/runtime/syscall_windows_test.go
//
// Without cgo can only use API calls, and callbacks of primitive types.
func doCallback() {
	f := func(p uintptr) uintptr {
		log.Println("non-Go thread callback")
		return p
	}

	createThreadProc := kernel32.MustFindProc("CreateThread")

	r, _, err := createThreadProc.Call(0, 0, windows.NewCallback(f), 123, 0, 0)
	if r == 0 {
		log.Fatalf("CreateThead failed, %v", err)
	}

	h := windows.Handle(r)
	defer windows.CloseHandle(h)

	switch s, err := windows.WaitForSingleObject(h, 1000); syscall.Errno(s) {
	case windows.WAIT_OBJECT_0:
		log.Println("successfully waited for thread")
	case windows.WAIT_TIMEOUT:
		log.Fatalf("timeout waiting for thread to exit")
	case windows.WAIT_FAILED:
		log.Fatalf("WaitForSingleObject failed, %v", err)
	default:
		log.Fatalf("WaitForSingleObject returns unexpected value, %v", s)
	}
}
