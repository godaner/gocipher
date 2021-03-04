package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
)

var DESDec = cli.Command{
	Name:      "desdec",
	Usage:     "decrypt by des",
	UsageText: "Usage: gocipher desdec [options...]",
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
			Name:     "key",
			Usage:    "key.",
			Required: false,
		},
		cli.StringFlag{
			Name:     "keyfile",
			Usage:    "key file.",
			Required: false,
		},
		cli.StringFlag{
			Name:     "o",
			Usage:    "write the plaintext to this file.",
			Required: false,
		},
		cli.StringFlag{
			Name:     "m",
			Usage:    "des mode, option is: cbc, triple.",
			Required: false,
			Value:    "cbc",
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
		text, textfile, key, keyfile, o, m, base64, d :=
			context.String("text"),
			context.String("textfile"),
			context.String("key"),
			context.String("keyfile"),
			context.String("o"),
			context.String("m"),
			context.String("base64"),
			context.Bool("d")
		logger.SetDebug(d)
		if key == "" && keyfile == "" {
			return ErrNeedPrivateKey
		}
		if key == "" {
			pbs, err := ioutil.ReadFile(keyfile)
			if err != nil {
				return err
			}
			key = string(pbs)
		}
		logger.Infof("Read key success, key is: %v", key)
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
		plainText, err := makeDESCipher(m).decrypt([]byte(text), []byte(key))
		if err != nil {
			return err
		}
		logger.Infof("Use key decrypt success, plaintext is: %v", string(plainText))
		if o != "" {
			if err := ioutil.WriteFile(o, []byte(string(plainText)+fmt.Sprintln()), 0777); err != nil {
				return err
			}
			logger.Infof("Save the plaintext to file success, file is: %v", o)
		}
		return nil
	},
}
