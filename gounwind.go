package gounwind

import (
	"unsafe"
)

//go:noinline

// Callers is a drop-in replacement for runtime.Callers that uses frame
// pointers for fast and simple stack unwinding.
func Callers(skip int, pcs []uintptr) int {
	return callers(skip, pcs)
}

//go:noinline
//go:nosplit
func callers(skip int, pcs []uintptr) int {
	fp := uintptr(unsafe.Pointer(&skip)) - 16
	i := 0
	for i < len(pcs) {
		pc := deref(fp + 8)
		if skip == 0 {
			pcs[i] = pc
			i++
		} else {
			skip--
		}
		fp = deref(fp)
		if fp == 0 {
			break
		}
	}
	return i
}

//go:nosplit
func deref(addr uintptr) uintptr {
	return *(*uintptr)(unsafe.Pointer(addr))
}
