package main

import (
	mlog "github.com/godaner/gocipher/log"
	"github.com/godaner/gocipher/log/logv1"
	"io/ioutil"
	"log"
	"os"
)

var logger mlog.Logger

func initLogger(debug bool) {
	logger = &logv1.LoggerV1{
		DebugWriter: ioutil.Discard,
		InfoWriter:  os.Stdout,
		WarnWriter:  os.Stdout,
		ErrorWriter: os.Stderr,
	}
	logger.SetDebug(debug)
	log.SetFlags(0)
}
