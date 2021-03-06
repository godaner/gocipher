package main

import "errors"

var (
	ErrNeedPrivateKey = errors.New("need private key")
	ErrNeedPublicKey  = errors.New("need public key")
	ErrNeedKey        = errors.New("need key")
	ErrNeedPlainText  = errors.New("need plain text")
	ErrNeedCipherText = errors.New("need cipher text")
)
