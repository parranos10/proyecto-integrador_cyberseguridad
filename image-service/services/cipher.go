package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

var ClaveSecreta = []byte("1234-es-123-clave-1234567-de-12-bytes!")

func EncryptAESCBC(plainText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plainText = padPKCS7(plainText, aes.BlockSize)

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return cipherText, nil
}

func DecryptAESCBC(cipherText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("el texto cifrado es demasiado corto")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	if len(cipherText)%aes.BlockSize != 0 {
		return nil, errors.New("el texto cifrado no es múltiplo del tamaño de bloque")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	return unpadPKCS7(cipherText)
}

func padPKCS7(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := make([]byte, padding)
	for i := range padText {
		padText[i] = byte(padding)
	}
	return append(src, padText...)
}

func unpadPKCS7(src []byte) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return nil, errors.New("bloque vacío")
	}
	unpadding := int(src[length-1])
	if unpadding > length {
		return nil, errors.New("padding incorrecto")
	}
	return src[:(length - unpadding)], nil
}
