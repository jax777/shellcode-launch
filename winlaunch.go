package main

import (
	"shellcode-launch/winshellcode"
)

func main() {
	sc := []byte("\xfc\x48\x83")
	winshellcode.Run(sc)
}
