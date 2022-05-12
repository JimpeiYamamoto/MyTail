package piscine

import (
	"os"
)

var ExitStatus int

const bufsize int = 32

type Tail struct {
	argv  []string
	fcnt  int
	farr  []string
	isOpt bool
	c     int
}

func (t *Tail) init() {
	var isAtoi, err bool

	len := SliceCount(t.argv)
	for i, s := range t.argv {
		if s == "-c" {
			t.isOpt = true
			if i+1 >= len {
				msg := "ztail: Option requires an argument -- 'c'\n"
				os.Stderr.WriteString(msg)
				os.Exit(1)
			}
			isAtoi = true
		} else if isAtoi == true {
			t.c, err = Atoi(s)
			if err {
				msg := "ztail: Invalid bytes -- `"
				os.Stderr.WriteString(msg + s + "'\n")
				os.Exit(1)
			}
			isAtoi = false
		} else {
			t.farr = append(t.farr, s)
			t.fcnt++
		}
	}
}

func (t Tail) putFiles() {
	for i, v := range t.farr {
		f, err := os.Open(v)
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			ExitStatus = 1
			continue
		}
		defer f.Close()
		if i != 0 && err == nil {
			os.Stderr.WriteString("\n")
		}
		if t.fcnt != 1 {
			os.Stdout.WriteString("==> " + v + " <==\n")
		}
		buf := make([]byte, bufsize)
		if t.isOpt {
			putLastBytes(f, t.c, buf)
		} else {
			putLastTenln(f, buf)
		}
	}
}

func Ztail(argv []string) {
	t := &Tail{argv: argv}
	t.init()
	t.putFiles()
}
