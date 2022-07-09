package main

import "testing"

func aaa(n int) int {
	for n > 0 {
		n--
	}
	return n
}

// benchmark 函数一般以 Benchmark 开头
// benchmark 的 case 一般会跑b.N次，而且每次执行都会如此
// 在执行过程中会根据实际 case 的执行时间是否稳定会增加 b.N 的次数以达到稳态
func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		//Print1to20()
		aaa(n)
	}
}

/**
 go test -bench=.
goos: darwin
goarch: amd64
pkg: go-learning/go-test/02
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkAll-16         179587407                6.753 ns/op
PASS
ok      go-learning/go-test/02  1.896s
*/
