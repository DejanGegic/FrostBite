package tools

import (
	_ "embed"
	"fmt"
	"os"
)

var pl = fmt.Println

func CopyPublicKeyFromAdmin() error {
	//copy key from admin/keys to main
	os.Mkdir("keys", 0755)
	pubKey, err := os.ReadFile("Admin/keys/public.key")
	if err != nil {
		return err
	}
	os.WriteFile("keys/public.key", pubKey, 0644)
	return nil
}
