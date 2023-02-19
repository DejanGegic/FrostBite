package encryption

import (
	"crypto/rand"
	"fmt"
	"os"
)

func DecryptUsingKeyFromFile() []byte {
	priv := ReadPrivKeyFromFile()
	encryptedData, err := os.ReadFile("keys/encrypted.key")
	ErrCheck(err)
	decryptedData := DecryptWithPrivateKey(encryptedData, BytesToPrivateKey(priv))
	return decryptedData
}

func EncryptUsingKeyFromFile(key []byte) []byte {
	pub := ReadPubKeyFromFile()
	encryptedData := EncryptWithPublicKey(key, BytesToPublicKey(pub))
	return encryptedData
}

func GenerateKeysAndWriteThemToFiles() {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	ErrCheck(err)

	privKey, pubKey := GenerateKeyPair(2048)

	os.Mkdir("keys", 0755)
	os.WriteFile("keys/private.key", PrivateKeyToBytes(privKey), 0644)
	os.WriteFile("keys/public.key", PublicKeyToBytes(pubKey), 0644)

	os.Mkdir("../keys", 0755)
	os.WriteFile("../keys/private.key", PrivateKeyToBytes(privKey), 0644)
	os.WriteFile("../keys/public.key", PublicKeyToBytes(pubKey), 0644)

	CheckIfFileKeysAreValid(key)
}

func CheckIfFileKeysAreValid(key []byte) {
	//read from file
	publicKeyFromFile, privateKeyFromFile := readKeysFromFIle()

	//convert to key
	pubKey := BytesToPublicKey(publicKeyFromFile)
	privKey := BytesToPrivateKey(privateKeyFromFile)

	// check if keys are valid
	encryptedData := EncryptWithPublicKey(key, pubKey)
	decryptedData := DecryptWithPrivateKey(encryptedData, privKey)
	if string(decryptedData) == string(key) {
		fmt.Println("Keys are valid")
	}
}

func readKeysFromFIle() (publicKey []byte, privateKey []byte) {
	publicKeyFromFile := ReadPubKeyFromFile()
	privateKeyFromFile := ReadPrivKeyFromFile()
	return publicKeyFromFile, privateKeyFromFile
}

func ReadPrivKeyFromFile() (privateKey []byte) {
	privateKeyFromFile, err := os.ReadFile("keys/private.key")
	ErrCheck(err)
	return privateKeyFromFile
}

func ReadPubKeyFromFile() (publicKey []byte) {
	publicKeyFromFile, err := os.ReadFile("keys/public.key")
	ErrCheck(err)
	return publicKeyFromFile
}
