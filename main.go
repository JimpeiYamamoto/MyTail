package main

import (
	"os"
	"piscine"
)

func main() {
	if piscine.SliceCount(os.Args) < 2 {
		return
	}
	piscine.Ztail(os.Args[1:])
	os.Exit(piscine.ExitStatus)
}
