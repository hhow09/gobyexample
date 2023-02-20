// performance-slice-append_test.go
package main

import "testing"

func BenchmarkGenByAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genByAppend()
	}
}

func BenchmarkGenByAppendCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genByAppendCap()
	}
}

func BenchmarkGenByAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genByAssign()
	}
}
