package main

import (
	"bytes"
	"crypto/rand"
	crsa "crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
)

var RSADec = cli.Command{
	Name:      "rsadec",
	Usage:     "decrypt by rsa",
	UsageText: "Usage: gocipher rsadec [options...]",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "text",
			Usage:    "cipher text.",
			Required: false,
		},
		cli.StringFlag{
			Name:     "textfile",
			Usage:    "cipher text file.",
			Required: false,
		},
		cli.StringFlag{
			Name:     "prikey",
			Usage:    "private key.",
			Required: false,
		},
		cli.StringFlag{
			Name:     "prikeyfile",
			Usage:    "private key file.",
			Required: false,
		},
		cli.StringFlag{
			Name:     "o",
			Usage:    "write the plaintext to this file.",
			Required: false,
		},
		cli.StringFlag{
			Name:     "base64",
			Usage:    "use base64 decode the ciphertext, option is: std, url, rawstd, rawurl.",
			Required: false,
		},
		cli.BoolFlag{
			Name:     "d",
			Usage:    "debug.",
			Required: false,
		},
	},
	Action: func(context *cli.Context) error {
		text, textfile, prikey, prikeyfile, o, base64, d :=
			context.String("text"),
			context.String("textfile"),
			context.String("prikey"),
			context.String("prikeyfile"),
			context.String("o"),
			context.String("base64"),
			context.Bool("d")
		logger.SetDebug(d)
		if prikey == "" && prikeyfile == "" {
			return ErrNeedPrivateKey
		}
		if prikey == "" {
			pbs, err := ioutil.ReadFile(prikeyfile)
			if err != nil {
				return err
			}
			prikey = string(pbs)
		}
		logger.Infof("Read private key success, private key is: %v", prikey)
		if text == "" && textfile == "" {
			return ErrNeedCipherText
		}
		if text == "" {
			bs, err := ioutil.ReadFile(textfile)
			if err != nil {
				return err
			}
			text = string(bs)
		}
		logger.Infof("Read cipher text success, cipher text is: %v", text)
		if base64 != "" {
			bs, err := makeBase64(base64).DecodeString(text)
			if err != nil {
				return err
			}
			text = string(bs)
		}
		plainText, err := decrypt([]byte(prikey), []byte(text))
		if err != nil {
			return err
		}
		logger.Infof("Use private key decrypt success, plaintext is: %v", string(plainText))
		if o != "" {
			if err := ioutil.WriteFile(o, []byte(string(plainText)+fmt.Sprintln()), 0777); err != nil {
				return err
			}
			logger.Infof("Save the plaintext to file success, file is: %v", o)
		}
		return nil
	},
}

// getRsaPrivateKey 获取RSA私钥
func getRsaPrivateKey(data []byte) (*crsa.PrivateKey, error) {
	block, _ := pem.Decode(data)
	private1KeyInterface, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return private1KeyInterface, nil
	}
	private8KeyInterface, err1 := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err1 == nil {
		return private8KeyInterface.(*crsa.PrivateKey), nil
	}
	return nil, errors.New(err.Error() + " and " + err1.Error())
}

// base64Decode base64解码
func base64Decode(data string) ([]byte, error) {
	base64er := base64.RawURLEncoding
	decode, err := base64er.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return decode, nil
}

// decrypt segment decrypt
func decrypt(privateKeyText, cipherText []byte) ([]byte, error) {
	privateKey, err := getRsaPrivateKey(privateKeyText)
	if err != nil {
		return nil, err
	}
	keySize := privateKey.N.BitLen() / 8
	cipherTextSize := len(cipherText)
	var offSet = 0
	var buffer = bytes.Buffer{}
	for offSet < cipherTextSize {
		endIndex := offSet + keySize
		if endIndex > cipherTextSize {
			endIndex = cipherTextSize
		}
		bytesOnce, err := crsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	return buffer.Bytes(), nil
}
