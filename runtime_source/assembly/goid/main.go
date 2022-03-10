package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"unsafe"
)

var offsetDictMap = map[string]int64{
	"go1.17": 152,
	"go1.10": 152,
	"go1.9":  152,
	"go1.8":  192,
}

var g_goid_offset = func() int64 {
	goversion := runtime.Version()
	for key, off := range offsetDictMap {
		if goversion == key || strings.HasPrefix(goversion, key) {
			return off
		}
	}
	panic("unsupport go verion:" + goversion)
}()

func GetGoid() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}

func getg() unsafe.Pointer

func GetGroutineId() int64 {
	g := getg()
	p := (*int64)(unsafe.Pointer(uintptr(g) + uintptr(g_goid_offset)))
	return *p
}

func main() {
	fmt.Println(GetGoid())
	fmt.Println(GetGroutineId())
}
