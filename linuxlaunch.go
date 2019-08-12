package main

import (
	"shellcode-launch/linuxshellcode"
)

func main() {
	sc := []byte("xaa\x51")
	linuxshellcode.Run(sc)
}
