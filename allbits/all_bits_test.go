package allbits

import (
	"fmt"
	"reflect"
	"testing"
)

var targetFuncs = []struct {
	name string
	body func(bitsLen int, callback func(bits []bool))
}{
	{"Bitwise", AllBitsBitwise},
	{"Recurisve", AllBitsRecurisve},
}

func TestAllBits(t *testing.T) {
	testCases := []struct {
		bitsLen int
		want    []string
	}{
		{0, []string{""}},
		{1, []string{"0", "1"}},
		{2, []string{"00", "01", "10", "11"}},
		{3, []string{"000", "001", "010", "011", "100", "101", "110", "111"}},
	}
	for _, targetFunc := range targetFuncs {
		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("%s(%d)", targetFunc.name, testCase.bitsLen), func(t *testing.T) {
				want := make([][]bool, len(testCase.want))
				for iBits, strBits := range testCase.want {
					bits := make([]bool, len(strBits))
					for iChar, charBit := range strBits {
						bits[iChar] = charBit == '1'
					}
					want[iBits] = bits
				}

				got := make([][]bool, 0)
				targetFunc.body(testCase.bitsLen, func(bits []bool) {
					bitsClone := make([]bool, len(bits))
					copy(bitsClone, bits)
					got = append(got, bitsClone)
				})

				if !reflect.DeepEqual(got, want) {
					t.Errorf("want: %v, got: %v", want, got)
				}
			})
		}
	}
}

func BenchmarkAllBits(b *testing.B) {
	for _, targetFunc := range targetFuncs {
		b.Run(targetFunc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				targetFunc.body(20, func(bits []bool) {
					sum := 0
					for _, bit := range bits {
						if bit {
							sum++
						}
					}
				})
			}
		})
	}
}
