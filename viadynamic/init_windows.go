package main

import "golang.org/x/sys/windows"

var (
	kernel32        = windows.MustLoadDLL("kernel32.dll")
	getModuleHandle = kernel32.MustFindProc("GetModuleHandleW")

	user32         = windows.MustLoadDLL("user32.dll")
	messageBoxProc = user32.MustFindProc("MessageBoxW")
)

func cleanup() {
	defer kernel32.Release()
	defer user32.Release()
}
