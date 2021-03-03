package main

import (
	"bytes"
	"crypto/rand"
	crsa "crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
)

var RSAEnc = cli.Command{
	Name:      "rsaenc",
	Usage:     "encrypt by rsa",
	UsageText: "Usage: gocipher rsaenc [options...]",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "text",
			Usage:    "plain text.",
			Required: false,
		},
		cli.StringFlag{
			Name:     "textfile",
			Usage:    "plain text file.",
			Required: false,
		},
		cli.StringFlag{
			Name:     "pubkey",
			Usage:    "public key.",
			Required: false,
		},
		cli.StringFlag{
			Name:     "pubkeyfile",
			Usage:    "public key file.",
			Required: false,
		},
		cli.StringFlag{
			Name:     "o",
			Usage:    "write the ciphertext to this file.",
			Required: false,
		},
		cli.StringFlag{
			Name:     "base64",
			Usage:    "use base64 encode the ciphertext, option is: std, url, rawstd, rawurl.",
			Required: false,
		},
		cli.BoolFlag{
			Name:     "d",
			Usage:    "debug.",
			Required: false,
		},
	},
	Action: func(context *cli.Context) error {
		text, textfile, pubkey, pubkeyfile, o, base64, d :=
			context.String("text"),
			context.String("textfile"),
			context.String("pubkey"),
			context.String("pubkeyfile"),
			context.String("o"),
			context.String("base64"),
			context.Bool("d")
		logger.SetDebug(d)
		if pubkey == "" && pubkeyfile == "" {
			return ErrNeedPublicKey
		}
		if pubkey == "" {
			pbs, err := ioutil.ReadFile(pubkeyfile)
			if err != nil {
				return err
			}
			pubkey = string(pbs)
		}
		logger.Infof("Read public key success, public key is: %v", pubkey)
		if text == "" && textfile == "" {
			return ErrNeedPlainText
		}
		if text == "" {
			bs, err := ioutil.ReadFile(textfile)
			if err != nil {
				return err
			}
			text = string(bs)
		}
		logger.Infof("Read cipher text success, plain text is: %v", text)
		cipherText, err := encrypt([]byte(pubkey), []byte(text))
		if err != nil {
			return err
		}
		if base64 != "" {
			cipherTextS := makeBase64(base64).EncodeToString(cipherText)
			cipherText = []byte(cipherTextS)
		}
		logger.Infof("Use public key encrypt success, ciphertext is: %v", string(cipherText))
		if o != "" {
			if err := ioutil.WriteFile(o, []byte(string(cipherText)+fmt.Sprintln()), 0777); err != nil {
				return err
			}
			logger.Infof("Save the ciphertext to file success, file is: %v", o)

		}
		return nil
	},
}

// getRsaPublicKey 获取RSA公钥
func getRsaPublicKey(data []byte) (*crsa.PublicKey, error) {
	block, _ := pem.Decode(data)
	var cert *x509.Certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPublicKey := cert.PublicKey.(*crsa.PublicKey)
	return rsaPublicKey, nil
}

// encrypt segment encrypt
func encrypt(publicKeyText, plaintext []byte) ([]byte, error) {
	publicKey, err := getRsaPublicKey(publicKeyText)
	if err != nil {
		return nil, err
	}
	keySize, srcSize := publicKey.N.BitLen()/8, len(plaintext)
	// keySize, srcSize := len(publicKey.N.Bytes()), len(src)
	// log.Println("密钥长度：", keySize, "\t明文长度：\t", srcSize)
	// 单次加密的长度需要减掉padding的长度，PKCS1为11
	offSet, once := 0, keySize-11
	buffer := bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + once
		if endIndex > srcSize {
			endIndex = srcSize
		}
		// 加密一部分
		bytesOnce, err := crsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}

	return buffer.Bytes(), nil
}
