package mshash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func MD5(text string) string {
	t := md5.New()
	if n, err := t.Write([]byte(text)); err != nil || n < 0 {
		return ""
	}
	return hex.EncodeToString(t.Sum(nil))
}

func SHA1(text string) string {
	t := sha1.New()
	if n, err := t.Write([]byte(text)); err != nil || n < 0 {
		return ""
	}
	return hex.EncodeToString(t.Sum(nil))
}

func SHA256(text string) string {
	t := sha256.New()
	if n, err := t.Write([]byte(text)); err != nil || n < 0 {
		return ""
	}
	return hex.EncodeToString(t.Sum(nil))
}

func SHA512(text string) string {
	t := sha512.New()
	if n, err := t.Write([]byte(text)); err != nil || n < 0 {
		return ""
	}
	return hex.EncodeToString(t.Sum(nil))
}
