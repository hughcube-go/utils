package mshash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func MD5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func SHA1(text string) string {
	t := sha1.New()
	t.Write([]byte(text))
	return hex.EncodeToString(t.Sum(nil))
}

func SHA256(text string) string {
	t := sha256.New()
	t.Write([]byte(text))
	return hex.EncodeToString(t.Sum(nil))
}

func SHA512(text string) string {
	t := sha512.New()
	t.Write([]byte(text))
	return hex.EncodeToString(t.Sum(nil))
}
