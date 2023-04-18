package foo_test

import (
	"example/foo"
	"testing"
)

func BenchmarkIncrementMutext(b *testing.B) {
	var a uint64 = 0
	for i := 0; i < b.N; i++ {
		foo.IncrementMutext(&a)
	}
}

func BenchmarkIncrementAtomic(b *testing.B) {
	var a uint64 = 0
	for i := 0; i < b.N; i++ {
		foo.IncrementAtomic(&a)
	}
}
