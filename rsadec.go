package main

import (
	"crypto/rand"
	crsa "crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
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
		cli.BoolFlag{
			Name:     "base64",
			Usage:    "use base64 decode the ciphertext.",
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
			context.String("private"),
			context.String("prikeyfile"),
			context.String("o"),
			context.Bool("base64"),
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
		if base64 {
			bs, err := base64Decode(text)
			if err != nil {
				return err
			}
			text = string(bs)
		}
		plainText, err := decrypt(text, prikey)
		if err != nil {
			return err
		}
		logger.Infof("Use private key decrypt success, plaintext is: %v", plainText)
		if o != "" {
			if err := ioutil.WriteFile(o, []byte(plainText+fmt.Sprintln()), 0777); err != nil {
				return err
			}
			logger.Infof("Save the plaintext to file success, file is: %v", o)
		}
		return nil
	},
}

func decrypt(ciphertext string, private string) (string, error) {
	prikey, err := getRsaPrivateKey([]byte(private))
	if err != nil {
		return "", err
	}
	plaintext, err := decrypto(prikey, []byte(ciphertext))
	if err != nil {
		return "", err
	}
	return string(plaintext[:]), nil
}

// getRsaPrivateKey 获取RSA私钥
func getRsaPrivateKey(data []byte) (*crsa.PrivateKey, error) {
	block, _ := pem.Decode(data)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaprivateKey := key
	return rsaprivateKey, nil
}

func decrypto(privatekey *crsa.PrivateKey, data []byte) ([]byte, error) {
	plainText, err := crsa.DecryptPKCS1v15(rand.Reader, privatekey, data)
	if err != nil {
		return nil, err
	}
	return plainText, nil
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
