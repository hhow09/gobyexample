$ go test -v -bench=. 

goos: darwin
goarch: arm64
pkg: github.com/hhow09/gobyexample/examples/performance-slice-append
BenchmarkGenByAppend
BenchmarkGenByAppend-8            563953              2148 ns/op
BenchmarkGenByAppendCap
BenchmarkGenByAppendCap-8        2595301               455.5 ns/op
BenchmarkGenByAssign
BenchmarkGenByAssign-8           4048084               294.9 ns/op
PASS
ok      github.com/hhow09/gobyexample/examples/performance-slice-append 5.483s