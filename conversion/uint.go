package conversion

import "fmt"

func Pow1(base, exponent int) int {
	ans := 1
	for i := 0; i < exponent; i++ {
		ans *= base
	}
	return ans
}

func Pow2(base, exponent int) int {
	if exponent < 0 {
		panic(fmt.Sprintf("exponent (%d) should not be a minus", exponent))
	}
	ans := 1
	for i := 0; i < exponent; i++ {
		ans *= base
	}
	return ans
}

func Pow3(base int, exponent uint) int {
	ans := 1
	for i := uint(0); i < exponent; i++ {
		ans *= base
	}
	return ans
}
