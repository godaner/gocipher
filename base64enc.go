package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
)

var BASE64Enc = cli.Command{
	Name:      "base64enc",
	Usage:     "encrypt by base64",
	UsageText: "Usage: gocipher base64enc [options...]",
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
		cli.StringFlag{
			Name:     "m",
			Usage:    "base64 mode, option is: std, url, rawstd, rawurl.",
			Required: true,
			Value:    "std",
		},
		cli.BoolFlag{
			Name:     "d",
			Usage:    "debug.",
			Required: false,
		},
	},
	Action: func(context *cli.Context) error {
		text, textfile, o, m, d :=
			context.String("text"),
			context.String("textfile"),
			context.String("o"),
			context.String("m"),
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
		logger.Infof("Read plain text success, plain text is: %v", text)
		cipherText := makeBase64(m).EncodeToString([]byte(text))
		logger.Infof("Use base64 encrypt success, ciphertext is: %v", cipherText)
		if o != "" {
			if err := ioutil.WriteFile(o, []byte(string(cipherText)+fmt.Sprintln()), 0777); err != nil {
				return err
			}
			logger.Infof("Save the ciphertext to file success, file is: %v", o)

		}
		return nil
	},
}
