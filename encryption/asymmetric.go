package encryption

import (
	"crypto/rand"
	"fmt"
	"os"
)

func DecryptUsingKeyFromFile() ([]byte, error) {
	priv, err := ReadPrivKeyFromFile()
	encryptedData, err := os.ReadFile("keys/encrypted.key")
	if err != nil {
		return nil, err
	}
	privateKeyBytes, err := BytesToPrivateKey(priv)
	if err != nil {
		return nil, err
	}
	decryptedData, err := DecryptWithPrivateKey(encryptedData, privateKeyBytes)
	if err != nil {
		return nil, err
	}
	return decryptedData, nil
}

func EncryptUsingKeyFromFile(key []byte) ([]byte, error) {
	pub, err := ReadPubKeyFromFile()
	if err != nil {
		return nil, err
	}
	rsaPubKeyBytes, err := BytesToPublicKey(pub)
	if err != nil {
		return nil, err
	}

	encryptedData, err := EncryptWithPublicKey(key, rsaPubKeyBytes)
	if err != nil {
		return nil, err
	}
	return encryptedData, err
}

func GenerateKeysAndWriteThemToFiles() error {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return err
	}

	privKey, pubKey, err := GenerateKeyPair(2048)
	if err != nil {
		return err
	}

	os.Mkdir("keys", 0755)
	pubKeyByte, err := PublicKeyToBytes(pubKey)
	privKeyByte := PrivateKeyToBytes(privKey)
	if err != nil {
		return err
	}
	os.WriteFile("keys/public.key", pubKeyByte, 0644)
	os.WriteFile("keys/private.key", privKeyByte, 0644)

	os.Mkdir("../keys", 0755)
	os.WriteFile("../keys/public.key", pubKeyByte, 0644)
	os.WriteFile("../keys/private.key", privKeyByte, 0644)

	err = CheckIfFileKeysAreValid(key)
	if err != nil {
		pl("Key validation failed")
		return err
	}
	return nil
}

func CheckIfFileKeysAreValid(key []byte) error {
	// read from file
	publicKeyFromFile, privateKeyFromFile, err := readKeysFromFIle()
	if err != nil {
		return err
	}

	// convert to key
	pubKey, err := BytesToPublicKey(publicKeyFromFile)
	if err != nil {
		return err
	}
	privKey, err := BytesToPrivateKey(privateKeyFromFile)

	// check if keys are valid
	encryptedData, err := EncryptWithPublicKey(key, pubKey)
	if err != nil {
		return err
	}
	decryptedData, err := DecryptWithPrivateKey(encryptedData, privKey)
	if err != nil {
		return err
	}
	if string(decryptedData) == string(key) {
		fmt.Println("Keys are valid")
	}
	return nil
}

func readKeysFromFIle() (publicKey []byte, privateKey []byte, err error) {
	publicKeyFromFile, err := ReadPubKeyFromFile()
	if err != nil {
		return nil, nil, err
	}
	privateKeyFromFile, err := ReadPrivKeyFromFile()
	if err != nil {
		return nil, nil, err
	}
	return publicKeyFromFile, privateKeyFromFile, nil
}

func ReadPrivKeyFromFile() (privateKey []byte, err error) {
	privateKeyFromFile, err := os.ReadFile("keys/private.key")
	if err != nil {
		return nil, err
	}
	return privateKeyFromFile, nil
}

func ReadPubKeyFromFile() (publicKey []byte, err error) {
	publicKeyFromFile, err := os.ReadFile("keys/public.key")
	if err != nil {
		return nil, err
	}
	return publicKeyFromFile, nil
}
