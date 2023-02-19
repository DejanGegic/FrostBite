package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"os"
)

func GenerateAES() []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	ErrCheck(err)
	return key
}
func DecryptFileAES(key []byte, filepath string) {

	block, err := aes.NewCipher(key)
	ErrCheck(err)
	gcm, err := cipher.NewGCM(block)
	ErrCheck(err)

	cipherText, err := os.ReadFile(filepath)
	ErrCheck(err)
	// //check if file ends in .enc
	// if filepath[len(filepath)-4:] != ".enc" {
	// 	return
	// }
	plainText, err := gcm.Open(nil, cipherText[:gcm.NonceSize()], cipherText[gcm.NonceSize():], nil)
	ErrCheck(err)
	originalFileName := filepath[:len(filepath)-4]
	os.WriteFile(originalFileName, plainText, 0644)
}

func EncryptFileAES(key []byte, filepath string) []byte {
	file := readFile(filepath)
	block, err := aes.NewCipher(key)
	ErrCheck(err)
	gcm, err := cipher.NewGCM(block)
	ErrCheck(err)

	nonce := make([]byte, gcm.NonceSize())
	ErrCheck(err)

	cipherText := gcm.Seal(nonce, nonce, file, nil)
	os.WriteFile(filepath+".enc", cipherText, 0644)
	return cipherText
}
func readFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	ErrCheck(err)
	return file
}
