package allbits

func AllBitsBitwise(bitsLen int, callback func(bits []bool)) {
	for pattern := 0; pattern < 1<<bitsLen; pattern++ {
		bits := make([]bool, bitsLen)
		for rpos := 0; rpos < bitsLen; rpos++ {
			bits[bitsLen-1-rpos] = pattern&(1<<rpos) > 0
		}

		callback(bits)
	}
}

func AllBitsRecurisve(bitsLen int, callback func(bits []bool)) {
	bits := make([]bool, bitsLen)
	var body func(pos int)
	body = func(pos int) {
		if pos == bitsLen {
			callback(bits)
			return
		}

		bits[pos] = false
		body(pos + 1)
		bits[pos] = true
		body(pos + 1)
	}
	body(0)
}
