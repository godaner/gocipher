package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func makeAESCipher(mode string) desCipher {
	switch mode {
	case "cbc":
		return &cbcAESCipher{}
	case "ecb":
		return &ecbAESCipher{}
	case "cfb":
		return &cfbAESCipher{}
	default:
		return &cbcAESCipher{}
	}
}

type cbcAESCipher struct {
}

func (c *cbcAESCipher) decrypt(cipherText, key []byte, ivDes ...byte) ([]byte, error) {
	return aesEncryptCBC(cipherText, key)
}
func (c *cbcAESCipher) encrypt(plainText, key []byte, ivDes ...byte) ([]byte, error) {
	return aesDecryptCBC(plainText, key)
}

type ecbAESCipher struct {
}

func (c *ecbAESCipher) decrypt(cipherText, key []byte, ivDes ...byte) ([]byte, error) {
	return aesEncryptECB(cipherText, key)
}
func (c *ecbAESCipher) encrypt(plainText, key []byte, ivDes ...byte) ([]byte, error) {
	return aesDecryptECB(plainText, key)
}

type cfbAESCipher struct {
}

func (c *cfbAESCipher) decrypt(cipherText, key []byte, ivDes ...byte) ([]byte, error) {
	return aesEncryptCFB(cipherText, key)
}
func (c *cfbAESCipher) encrypt(plainText, key []byte, ivDes ...byte) ([]byte, error) {
	return aesDecryptCFB(plainText, key)
}

// =================== CBC ======================
func aesEncryptCBC(origData []byte, key []byte) (encrypted []byte, err error) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted, nil
}
func aesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte, err error) {
	block, _ := aes.NewCipher(key)                              // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return decrypted, nil
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// =================== ECB ======================
func aesEncryptECB(origData []byte, key []byte) (encrypted []byte, err error) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted, nil
}
func aesDecryptECB(encrypted []byte, key []byte) (decrypted []byte, err error) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	bEnd := SearchByteSliceIndex(decrypted, 0)

	return decrypted[:bEnd], nil
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// []byte 字节切片 循环查找
func SearchByteSliceIndex(bSrc []byte, b byte) int {
	for i := 0; i < len(bSrc); i++ {
		if bSrc[i] == b {
			return i
		}
	}

	return -1
}

// =================== CFB ======================
func aesEncryptCFB(origData []byte, key []byte) (encrypted []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted, nil
}
func aesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte, err error) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		return nil, ErrCipherTextTooShort
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted, nil
}
