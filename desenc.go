package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
)

var DESEnc = cli.Command{
	Name:      "desenc",
	Usage:     "encrypt by des",
	UsageText: "Usage: gocipher desenc [options...]",
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
			Usage:    "write the ciphertext to this file.",
			Required: false,
		},
		cli.StringFlag{
			Name:     "base64",
			Usage:    "use base64 encode the ciphertext, option is: std, url, rawstd, rawurl.",
			Required: false,
		},
		cli.StringFlag{
			Name:     "m",
			Usage:    "des mode, option is: cbc, triple.",
			Required: false,
			Value:    "cbc",
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
			return ErrNeedKey
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
			return ErrNeedPlainText
		}
		if text == "" {
			bs, err := ioutil.ReadFile(textfile)
			if err != nil {
				return err
			}
			text = string(bs)
		}
		logger.Infof("Read plain text success, plain text is: %v", text)
		cipherText, err := makeDESCipher(m).encrypt([]byte(text), []byte(key))
		if err != nil {
			return err
		}
		if base64 != "" {
			cipherTextS := makeBase64(base64).EncodeToString(cipherText)
			cipherText = []byte(cipherTextS)
		}
		logger.Infof("Use key encrypt success, ciphertext is: %v", string(cipherText))
		if o != "" {
			if err := ioutil.WriteFile(o, []byte(string(cipherText)+fmt.Sprintln()), 0777); err != nil {
				return err
			}
			logger.Infof("Save the ciphertext to file success, file is: %v", o)

		}
		return nil
	},
}
