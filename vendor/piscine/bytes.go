package piscine

func countInBytes(buf []byte, b byte) int {
	var ret int
	for _, v := range buf {
		if v == b {
			ret++
		}
	}
	return ret
}

func searchInBytes(buf []byte, b byte) ([]byte, int) {
	for i, v := range buf {
		if v == b {
			return buf[i+1:], i + 1
		}
	}
	return buf, -1
}
