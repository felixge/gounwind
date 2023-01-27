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
	fp := regfp()
	i := 0
	for i < len(pcs) {
		pc := *(*uintptr)(unsafe.Pointer(fp + 8))
		if skip == 0 {
			pcs[i] = pc
			i++
		} else {
			skip--
		}
		fp = *(*uintptr)(unsafe.Pointer(fp))
		if fp == 0 {
			break
		}
	}
	return i
}

// regfp returns the frame pointer addr in the callers frame by
func regfp() uintptr
