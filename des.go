package main

import "github.com/wumansgy/goEncrypt"

func makeDESCipher(mode string) desCipher {
	switch mode {
	case "", "cbc":
		return &cbcDESCipher{}
	case "3", "triple":
		return &tripleDESCipher{}
	default:
		return &cbcDESCipher{}
	}
}

type desCipher interface {
	encrypt(plainText, key []byte, ivDes ...byte) ([]byte, error)
	decrypt(cipherText, key []byte, ivDes ...byte) ([]byte, error)
}
type tripleDESCipher struct {
}

func (c *tripleDESCipher) decrypt(cipherText, key []byte, ivDes ...byte) ([]byte, error) {
	return goEncrypt.TripleDesDecrypt(cipherText, key, ivDes...)
}

func (c *tripleDESCipher) encrypt(plainText, key []byte, ivDes ...byte) ([]byte, error) {
	return goEncrypt.TripleDesEncrypt(plainText, key, ivDes...)
}

type cbcDESCipher struct {
}

func (c *cbcDESCipher) decrypt(cipherText, key []byte, ivDes ...byte) ([]byte, error) {
	return goEncrypt.DesCbcDecrypt(cipherText, key, ivDes...)
}
func (c *cbcDESCipher) encrypt(plainText, key []byte, ivDes ...byte) ([]byte, error) {
	return goEncrypt.DesCbcEncrypt(plainText, key, ivDes...)
}
