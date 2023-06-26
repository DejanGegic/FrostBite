package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
)

var pl = fmt.Println

func GenerateAES() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		pl("Error generating AES")
		return nil, err
	}
	return key, nil
}
func DecryptFileAES(key []byte, filepath string) error {

	block, err := aes.NewCipher(key)
	if err != nil {
		pl("Error generating AES cipher")
		return err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		pl("Error generating CGm block")
		return err
	}

	cipherText, err := os.ReadFile(filepath)
	if err != nil {
		pl("Error reading file")
		return err
	}

	plainText, err := gcm.Open(nil, cipherText[:gcm.NonceSize()], cipherText[gcm.NonceSize():], nil)
	if err != nil {
		pl("Error opening gcm")
		return err
	}
	originalFileName := filepath[:len(filepath)-4]
	os.WriteFile(originalFileName, plainText, 0644)

	return nil
}

func EncryptFileAES(key []byte, filepath string) ([]byte, error) {
	file, err := readFile(filepath)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nonce, nonce, file, nil)
	os.WriteFile(filepath+".enc", cipherText, 0644)
	return cipherText, nil
}
func readFile(filePath string) ([]byte, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		pl("Error reading file")
		return nil, err
	}
	//close file
	return file, nil
}
