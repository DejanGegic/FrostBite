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

	var mode []string
	// take input from command line of flag --decrypt or --generate
	if len(os.Args) > 1 {
		mode = os.Args[1:]
	}
	pl("Mode: ", mode)
	for _, m := range mode {
		switch m {
		case "--generate":
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
		case "--decrypt":
			key, err = os.ReadFile("keys/encrypted.key")
			ErrCheck(err)
			decrypted := enc.DecryptUsingKeyFromFile()
			os.WriteFile("keys/decrypted.key", []byte(decrypted), 0644)
			readFromFileKey, _ := os.ReadFile("keys/decrypted.key")
			pl("Decrypted key: ", string(readFromFileKey))

		default:
			pl("Invalid flag")
		}
	}

}
