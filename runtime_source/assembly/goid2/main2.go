package main

import (
	"unsafe"
)

func getg() int64

var g_goid_offset uintptr = 152

func GetGroutineId() int64 {
	g := getg()
	p := (*int64)(unsafe.Pointer(uintptr(g) + g_goid_offset))
	return *p
}
