package tools

import (
	_ "embed"
	"fmt"
	"os"
)

var pl = fmt.Println

func ErrCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func CopyPublicKeyFromAdmin() {
	//copy key from admin/keys to main
	os.Mkdir("keys", 0755)
	pubKey, err := os.ReadFile("Admin/keys/public.key")
	ErrCheck(err)
	os.WriteFile("keys/public.key", pubKey, 0644)
}
