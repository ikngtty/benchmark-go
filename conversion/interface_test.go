package conversion

import (
	"testing"
)

func BenchmarkList(b *testing.B) {
	const addCount = 100_000
	const expectedTotal = addCount * (addCount - 1) / 2

	b.Run("IntList", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			list := NewIntList()
			for j := 0; j < addCount; j++ {
				list.Add(j)
			}

			total := 0
			list.Each(func(elem int) {
				total += elem
			})

			if total != expectedTotal {
				b.Errorf("got: %d, want:%d", total, expectedTotal)
			}
		}
	})
	b.Run("AnyList", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			list := NewAnyList()
			for j := 0; j < addCount; j++ {
				list.Add(j)
			}

			total := 0
			list.Each(func(elem interface{}) {
				total += elem.(int)
			})

			if total != expectedTotal {
				b.Errorf("got: %d, want:%d", total, expectedTotal)
			}
		}
	})
}
