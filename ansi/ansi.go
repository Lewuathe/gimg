package ansi

import (
	"fmt"
	// TODO: Support PNG
	//"image/png"
)

const escape = "\x1b"

func To256(c uint32) uint32 {
	ret := (float32(c) / 65536.0) * 256
	return uint32(ret)
}

// Print in 256-xterm color by using escape sequences
func Print(r, g, b uint32) {
	fmt.Printf("%s[7m%s[38;2;%d;%d;%dm  ", escape, escape, r, g, b)
}
