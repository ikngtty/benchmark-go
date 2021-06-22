# eucdiv

"eucdiv" means the Euclidean divison.

```
goos: darwin
goarch: amd64
pkg: github.com/ikngtty/benchmark-go/eucdiv
cpu: Intel(R) Core(TM) i7-1060NG7 CPU @ 1.20GHz
BenchmarkEucRem/EucRem1(00000000007,00000000003)-8         	371334560	         3.001 ns/op	       0 B/op	       0 allocs/op
BenchmarkEucRem/EucRem1(-0000000007,00000000003)-8         	398011891	         3.081 ns/op	       0 B/op	       0 allocs/op
BenchmarkEucRem/EucRem1(09999999999,01000000007)-8         	403263966	         2.996 ns/op	       0 B/op	       0 allocs/op
BenchmarkEucRem/EucRem1(-9999999999,01000000007)-8         	381840141	         3.158 ns/op	       0 B/op	       0 allocs/op
BenchmarkEucRem/EucRem2(00000000007,00000000003)-8         	209458334	         5.736 ns/op	       0 B/op	       0 allocs/op
BenchmarkEucRem/EucRem2(-0000000007,00000000003)-8         	208147119	         5.815 ns/op	       0 B/op	       0 allocs/op
BenchmarkEucRem/EucRem2(09999999999,01000000007)-8         	206962057	         5.818 ns/op	       0 B/op	       0 allocs/op
BenchmarkEucRem/EucRem2(-9999999999,01000000007)-8         	206765756	         5.774 ns/op	       0 B/op	       0 allocs/op
```
