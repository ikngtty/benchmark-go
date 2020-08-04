package uintconv

import (
	"fmt"
	"testing"
)

var powFuncs = []struct {
	name string
	body func(int, int) int
}{
	{"Pow1", Pow1},
	{"Pow2", Pow2},
	{"Pow3", func(base, exponent int) int {
		return Pow3(base, uint(exponent))
	}},
}

func TestPow(t *testing.T) {
	cases := []struct {
		base     int
		exponent int
		want     int
	}{
		{5, 0, 1},
		{5, 1, 5},
		{5, 2, 25},
		{5, 3, 125},
	}
	for _, f := range powFuncs {
		for _, c := range cases {
			caseName := fmt.Sprintf("%s(%d,%d)", f.name, c.base, c.exponent)
			t.Run(caseName, func(t *testing.T) {
				got := f.body(c.base, c.exponent)
				if got != c.want {
					t.Errorf("got: %d, want: %d", got, c.want)
				}
			})
		}
	}
}

func BenchmarkPow(b *testing.B) {
	for _, f := range powFuncs {
		b.Run(f.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f.body(2, 10)
			}
		})
	}
}
