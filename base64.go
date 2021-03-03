package main

import "encoding/base64"

// makeBase64 option is: std, url, rawstd, rawurl
func makeBase64(w string) *base64.Encoding {
	switch w {
	case "std":
		return base64.StdEncoding
	case "url":
		return base64.URLEncoding
	case "rawstd":
		return base64.RawStdEncoding
	case "rawurl":
		return base64.RawURLEncoding
	default:
		return base64.StdEncoding
	}
}
