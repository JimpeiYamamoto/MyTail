package piscine

func SliceCount(arr []string) int {
	var ret int

	for range arr {
		ret++
	}
	return ret
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func Atoi(s string) (int, bool) {
	var ret int
	var isMinus bool

	if s == "" {
		return 0, true
	}
	runes := []rune(s)
	if runes[0] == '+' || runes[0] == '-' {
		if runes[0] == '-' {
			isMinus = true
		}
		runes = runes[1:]
	}
	for _, v := range runes {
		if v < '0' || v > '9' {
			return 0, true
		}
	}
	for _, v := range runes {
		if isMinus {
			if ret >= 10*ret-int(v-'0') {
				ret = 10*ret - int(v-'0')
			} else {
				ret = int(^uint(0) >> 1)
				return ret + 1, true
			}
		} else {
			if ret <= 10*ret+int(v-'0') {
				ret = 10*ret + int(v-'0')
			} else {
				return int(^uint(0) >> 1), true
			}
		}
	}
	return ret, false
}
