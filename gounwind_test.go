package gounwind

import (
	"runtime"
	"testing"
)

func TestCallers(t *testing.T) {
	for _, m := range []method{runtimeCallers, gounwindCallers} {
		t.Run(string(m), func(t *testing.T) {
			for _, fn := range resolvePCs(testCallers(m)) {
				t.Logf("%s", fn)
			}
		})
	}
}

func BenchmarkUnwind(b *testing.B) {
	for _, m := range []method{runtimeCallers, gounwindCallers} {
		b.Run(string(m), func(b *testing.B) {
			bench(b, m, 16)
		})
	}
}

//go:noinline
func bench(b *testing.B, m method, depth int) {
	pcs := make([]uintptr, depth+10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var n int
		switch m {
		case runtimeCallers:
			n = runtime.Callers(1, pcs)
		case gounwindCallers:
			n = Callers(1, pcs)
		}
		if n > depth {
			panic("bad")
		} else if n < depth {
			bench(b, m, depth)
			break
		}
	}
}

type method string

const (
	runtimeCallers  method = "runtime"
	gounwindCallers method = "gounwind"
)

func resolvePCs(pcs []uintptr) []string {
	fns := make([]string, 0, len(pcs))
	frames := runtime.CallersFrames(pcs)
	for {
		frame, more := frames.Next()
		fns = append(fns, frame.Function)
		if !more {
			break
		}
	}
	return fns
}

//go:noinline
func testCallers(m method) []uintptr {
	return testCallersA(m, 5)
}

//go:noinline
func testCallersA(m method, i int) []uintptr {
	return testCallersB(m, i, 9)
}

//go:noinline
func testCallersB(m method, i, j int) []uintptr {
	pcs := make([]uintptr, 32)
	switch m {
	case runtimeCallers:
		return pcs[0:runtime.Callers(1, pcs)]
	case gounwindCallers:
		return pcs[0:Callers(1, pcs)]
	}
	panic("unreachable")
}
