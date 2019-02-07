package factorial

import (
	"testing"
)

const (
	benchN = 1000
)

func BenchmarkRecur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Recur(benchN)
	}
}

func BenchmarkIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iter(benchN)
	}
}
