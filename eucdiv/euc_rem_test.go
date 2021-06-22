package eucrem

import (
	"fmt"
	"testing"
)

var targetFuncs = []struct {
	name string
	body func(divident, dividor int) int
}{
	{"EucRem1", EucRem1},
	{"EucRem2", EucRem2},
}

func TestEucRem(t *testing.T) {
	cases := []struct {
		divident int
		dividor  int
		want     int
	}{
		{0, 3, 0}, {0, -3, 0},
		{7, 1, 0}, {-7, 1, 0},
		{7, 3, 1}, {-7, 3, 2},
		{7, 10, 7}, {-7, 10, 3},
	}

	for _, targetFunc := range targetFuncs {
		for _, c := range cases {
			caseName := fmt.Sprintf("%s(%d,%d)", targetFunc.name, c.divident, c.dividor)
			t.Run(caseName, func(t *testing.T) {
				got := targetFunc.body(c.divident, c.dividor)
				if got != c.want {
					t.Errorf("want: %d, got: %d", c.want, got)
				}
			})
		}
	}
}

func BenchmarkEucRem(b *testing.B) {
	cases := []struct {
		divident int
		dividor  int
	}{
		{7, 3}, {-7, 3},
		{9_999_999_999, 1_000_000_007}, {-9_999_999_999, 1_000_000_007},
	}

	for _, targetFunc := range targetFuncs {
		for _, c := range cases {
			caseName := fmt.Sprintf("%s(%011d,%011d)", targetFunc.name, c.divident, c.dividor)
			b.Run(caseName, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					targetFunc.body(c.divident, c.dividor)
				}
			})
		}
	}
}
