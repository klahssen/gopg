package factorial

import (
	"testing"

	"github.com/klahssen/tester"
)

const (
	benchN = 1000
)

func BenchmarkRecur(b *testing.B) {
	//t0 := time.Now()
	//fmt.Printf("benchmark recur\n")
	for i := 0; i < b.N; i++ {
		//t0 = time.Now()
		Recur(benchN)
		//fmt.Printf("done in %s\n", time.Since(t0))
	}
}

func BenchmarkIter(b *testing.B) {
	//t0 := time.Now()
	//fmt.Printf("benchmark iter\n")
	for i := 0; i < b.N; i++ {
		//t0 = time.Now()
		Iter(benchN)
		//fmt.Printf("done in %s\n", time.Since(t0))
	}
}

func TestRecur(t *testing.T) {
	te := tester.NewT(t)
	tests := []struct {
		n   int
		res int
	}{
		{5, 5 * 4 * 3 * 2 * 1},
	}
	for ind, test := range tests {
		te.DeepEqual(ind, "res", test.res, Recur(test.n))
	}
}
