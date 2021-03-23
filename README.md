# gounwind

gounwind is a tiny Go library that acts as a drop-in replacement for [`runtime.Callers()`](https://golang.org/pkg/runtime/#Callers). It exists to show how simple and fast stack unwinding can theoretically be when using frame pointers.

Compared to `runtime.Callers()`, gounwind is:

- ~50 faster
- [~25 lines of code](./gounwind.go) vs [thousands](https://github.com/golang/go/blob/go1.16.2/src/runtime/traceback.go#L76-L559)
- Totally unsafe for production use
- Unable to recognize inlined functions
- Only works on 64 bit platforms [where frame pointers are enabled](https://github.com/golang/go/blob/go1.16.2/src/runtime/runtime2.go#L1108)

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

## More Information

The [The Go low-level calling convention on x86-64](https://dr-knz.net/go-calling-convention-x86-64.html) article as well as [this video](https://www.youtube.com/watch?v=PrDsGldP1Q0) were incredibly useful to me while trying to figure this out.

Go includes frame pointers by default [since Go 1.7](https://github.com/golang/go/issues/15840).

There is [ongoing work](https://github.com/golang/go/issues/16638) to implement frame pointer unwinding for the Go core.
