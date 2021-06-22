package eucrem

func EucRem1(divident, dividor int) int {
	// suppose that dividor != 0
	rem := divident % dividor
	if rem < 0 {
		rem += Abs(dividor)
	}
	return rem
}

func EucRem2(divident, dividor int) int {
	// suppose that dividor > 0
	return (divident%dividor + dividor) % dividor
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
