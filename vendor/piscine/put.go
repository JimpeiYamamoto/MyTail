package piscine

import (
	"os"
)

func putLastBytes(f *os.File, c int, buf []byte) {
	start := fileSize(f) - c
	if start < 0 {
		start = 0
	}
	for c > 0 {
		n := myReadAt(f, buf, start)
		os.Stdout.Write(buf[:n])
		start += n
		c -= bufsize
		if n == 0 {
			break
		}
	}
}

func putLastTenln(f *os.File, buf []byte) {
	size := fileSize(f)
	at := size - bufsize
	line := 0
	for {
		if at < 0 {
			n := myReadAt(f, buf, 0)
			line += countInBytes(buf[:n], '\n')
			break
		} else {
			n := myReadAt(f, buf, at)
			line += countInBytes(buf[:n], '\n')
			if line > 10 {
				break
			}
		}
		at -= bufsize
	}
	if line > 10 {
		at = max(at, 0)
		n := myReadAt(f, buf, at)
		buf = buf[:n]
		for line > 10 {
			i := 0
			buf, i = searchInBytes(buf, '\n')
			at += i
			line--
		}
	} else {
		at = 0
	}
	for at < size {
		n := myReadAt(f, buf, at)
		os.Stdout.Write(buf[:n])
		at += n
	}
}
