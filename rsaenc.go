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
		cli.BoolFlag{
			Name:     "base64",
			Usage:    "use base64 encode the ciphertext.",
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
			context.Bool("base64"),
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
		cipherText, err := encrypt(text, pubkey)
		if err != nil {
			return err
		}
		if base64 {
			cipherText = base64Encode([]byte(cipherText))
		}
		logger.Infof("Use public key encrypt success, ciphertext is: %v", cipherText)
		if o != "" {
			if err := ioutil.WriteFile(o, []byte(cipherText+fmt.Sprintln()), 0777); err != nil {
				return err
			}
			logger.Infof("Save the ciphertext to file success, file is: %v", o)

		}
		return nil
	},
}

func encrypt(plaintext string, public string) (string, error) {
	key, err := getRsaPublicKey([]byte(public))
	if err != nil {
		return "", err
	}
	ciphertext, err := encrypto(key, []byte(plaintext))
	if err != nil {
		return "", err
	}
	return string(ciphertext[:]), nil
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

func encrypto(publickey *crsa.PublicKey, data []byte) ([]byte, error) {
	ciphertext, err := crsa.EncryptPKCS1v15(rand.Reader, publickey, data)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// base64Encode base64编码
func base64Encode(data []byte) string {
	base64er := base64.RawURLEncoding
	return base64er.EncodeToString(data)
}
