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
	app.Usage = "Gocipher is a cross platform command line tool for encryption and decryption, including RSA, DES."
	app.Version = "v1.0.3"
	app.Commands = []cli.Command{
		RSAEnc,
		RSADec,
		DESEnc,
		DESDec,
	}
	if err := app.Run(os.Args); err != nil {
		logger.Error(err)
		os.Exit(0)
	}
}
