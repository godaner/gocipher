package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
)

var Sha256 = cli.Command{
	Name:      "sha256",
	Usage:     "sha256",
	UsageText: "Usage: gocipher sha256 [options...]",
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
			Name:     "o",
			Usage:    "write the ciphertext to this file.",
			Required: false,
		},
		cli.BoolFlag{
			Name:     "d",
			Usage:    "debug.",
			Required: false,
		},
	},
	Action: func(context *cli.Context) error {
		text, textfile, o, d :=
			context.String("text"),
			context.String("textfile"),
			context.String("o"),
			context.Bool("d")
		logger.SetDebug(d)
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
		// logger.Infof("Read plain text success, plain text is: %v", text)
		cipherText := fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
		logger.Infof("Use sha256 encrypt success, ciphertext is: %v", cipherText)
		if o != "" {
			if err := ioutil.WriteFile(o, []byte(cipherText+fmt.Sprintln()), 0777); err != nil {
				return err
			}
			logger.Infof("Save the ciphertext to file success, file is: %v", o)

		}
		return nil
	},
}
