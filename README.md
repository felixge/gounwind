# gounwind

gounwind is a tiny Go library that acts as a drop-in replacement for [`runtime.Callers()`](https://golang.org/pkg/runtime/#Callers). It's highly experimental and probably only works on x86-64. Do not use it in production.

Compared to `runtime.Callers()`, gounwind is:

- 55x faster
- [~25 lines of code](./unwind.go) vs [thousands](https://github.com/golang/go/blob/go1.16.2/src/runtime/traceback.go#L76-L559)
- Totally unsafe for production use
- Unable to recognize inlined functions
- Only works on [64 bit platforms](https://github.com/golang/go/blob/go1.16.2/src/runtime/runtime2.go#L1108) where frame pointers are enabled

## Benchmark

The benchmark below shows the performance for unwinding a stack that has 16 frames. The numbers are from my macOS machine and Docker for Linux gives me very similar results.

```
goos: darwin
goarch: amd64
pkg: github.com/felixge/gounwind
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkUnwind
BenchmarkUnwind/runtime
BenchmarkUnwind/runtime-12         	1281306	      934.7 ns/op
BenchmarkUnwind/gounwind
BenchmarkUnwind/gounwind-12        	65443237	       17.69 ns/op
```
