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
	i := 0
	frame := (*frame)(unsafe.Pointer(regfp()))
	for i < len(pcs) {
		if skip == 0 {
			pcs[i] = frame.retpc
			i++
		} else {
			skip--
		}
		if frame.pointer == nil {
			break
		}
		frame = frame.pointer
	}
	return i
}

type frame struct {
	pointer *frame
	retpc   uintptr
}

// regfp returns the frame pointer addr in the callers frame by
func regfp() uintptr
