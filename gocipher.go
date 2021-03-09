package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	initLogger(false)
	cli.HelpFlag = cli.BoolFlag{
		Name:  "help",
		Usage: "show help",
	}
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version",
		Usage: "print the version",
	}
	app := cli.NewApp()
	app.Writer = os.Stdout
	app.ErrWriter = os.Stderr
	app.Name = "gocipher"
	app.HelpName = "gocipher"
	app.Usage = "Gocipher is a cross platform command line tool for encryption and decryption, including RSA, DES, BASE64, MD5, SHA256."
	app.Version = "v1.0.4"
	app.Commands = []cli.Command{
		RSAEnc,
		RSADec,
		DESEnc,
		DESDec,
		BASE64Enc,
		BASE64Dec,
		Md5,
		Sha256,
	}
	if err := app.Run(os.Args); err != nil {
		logger.Error(err)
		os.Exit(0)
	}
}
