package main

import (
	"crypto/rand"
	"fmt"
	"os"

	enc "admin/encryption"
)

func ErrCheck(err error) {
	if err != nil {
		panic(err)
	}
}

var pl = fmt.Println

func main() {

	// GenerateKeysAndWriteThemToFiles()

	key := make([]byte, 32)
	_, err := rand.Read(key)
	ErrCheck(err)

	// encryptedData := enc.EncryptUsingKeyFromFile(key)
	// os.WriteFile("keys/encrypted.key", encryptedData, 0644)

	// decryptedData := DecryptUsingKeyFromFile(encryptedData, err)
	// os.WriteFile("keys/decrypted.key", []byte(hex.EncodeToString(decryptedData)), 0644)

	var mode string
	// take input from command line of flag --decrypt or --generate
	mode = ChoseMode()

	pl("Mode: ", mode)
	switch mode {
	case "generate":
		//confirm if user wants to generate new keys, if yes, then generate
		generate := ""
		fmt.Print("Are you sure you want to generate new keys? (y/N)  ")
		fmt.Scanln(&generate)
		if generate == "y" {
			enc.GenerateKeysAndWriteThemToFiles()
		} else {
			pl("Exiting...")
			os.Exit(0)
		}
	case "decrypt":
		key, err = os.ReadFile("keys/encrypted.key")
		if err != nil {
			fmt.Println("Please put encrypted key in keys/encrypted.key. Exiting...")
			break
		}
		decrypted := enc.DecryptUsingKeyFromFile()
		os.WriteFile("keys/decrypted.key", []byte(decrypted), 0644)
		readFromFileKey, _ := os.ReadFile("keys/decrypted.key")
		pl("Decrypted key: ", string(readFromFileKey))

	default:
		pl("Invalid flag")
	}

}

func ChoseMode() string {
	// Mode depends on if "keys" dir exists and if it contains "public.key" and "private.key" files
	// If "keys" dir exists and contains "public.key" and "private.key" files, then mode is "decrypt"
	// If "keys" dir exists and does not contain "public.key" and "private.key" files, then mode is "generate"
	// If "keys" dir does not exist, then mode is "generate"
	_, err := os.Stat("keys")
	if os.IsNotExist(err) {
		return "generate"
	} else {
		pl("Keys dir exists")
		_, errPub := os.Stat("keys/public.key")
		_, errPriv := os.Stat("keys/private.key")
		if os.IsNotExist(errPub) || os.IsNotExist(errPriv) {
			return "generate"
		} else {
			return "decrypt"
		}
	}
}
