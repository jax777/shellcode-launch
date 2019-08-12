package main

import (
	"shellcode-launch/linuxshellcode"
)

func main() {
	sc := []byte("xaa\x51\xc3.........")
	winshellcode.Run(sc)
}
