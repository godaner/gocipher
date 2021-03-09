package main

import "github.com/wumansgy/goEncrypt"

func makeAESCipher(mode string) desCipher {
	switch mode {
	case "cbc":
		return &cbcAESCipher{}
	case "ctr":
		return &crtAESCipher{}
	default:
		return &cbcAESCipher{}
	}
}

type cbcAESCipher struct {
}

func (c *cbcAESCipher) decrypt(cipherText, key []byte, ivDes ...byte) ([]byte, error) {
	return goEncrypt.AesCbcEncrypt(cipherText, key, ivDes...)
}
func (c *cbcAESCipher) encrypt(plainText, key []byte, ivDes ...byte) ([]byte, error) {
	return goEncrypt.AesCbcDecrypt(plainText, key, ivDes...)
}

type crtAESCipher struct {
}

func (c *crtAESCipher) decrypt(cipherText, key []byte, ivDes ...byte) ([]byte, error) {
	return goEncrypt.AesCtrEncrypt(cipherText, key, ivDes...)
}
func (c *crtAESCipher) encrypt(plainText, key []byte, ivDes ...byte) ([]byte, error) {
	return goEncrypt.AesCtrDecrypt(plainText, key, ivDes...)
}
