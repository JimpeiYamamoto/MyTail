package piscine

import "os"

func fileSize(f *os.File) int {
	st, err := os.Stat(f.Name())
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
	return int(st.Size())
}

func myReadAt(f *os.File, buf []byte, start int) int {
	n, err := f.ReadAt(buf, int64(start))
	if err != nil && err.Error() != "EOF" {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
	return n
}
