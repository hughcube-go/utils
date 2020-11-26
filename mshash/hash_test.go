package mshash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func test_Getitems() []string {
	items := []string{"1", "你好", "hello", ""}

	return items
}

func Test_MD5(t *testing.T) {
	a := assert.New(t)

	for _, v := range test_Getitems() {

		h := md5.New()
		h.Write([]byte(v))
		hv := hex.EncodeToString(h.Sum(nil))

		a.Equal(hv, MD5(v))
	}
}

func Test_SHA1(t *testing.T) {
	a := assert.New(t)

	for _, v := range test_Getitems() {

		t := sha1.New()
		t.Write([]byte(v))
		hv := hex.EncodeToString(t.Sum(nil))

		a.Equal(hv, SHA1(v))
	}
}

func Test_SHA256(t *testing.T) {
	a := assert.New(t)

	for _, v := range test_Getitems() {

		t := sha256.New()
		t.Write([]byte(v))
		hv := hex.EncodeToString(t.Sum(nil))

		a.Equal(hv, SHA256(v))
	}
}


func Test_SHA512(t *testing.T) {
	a := assert.New(t)

	for _, v := range test_Getitems() {

		t := sha512.New()
		t.Write([]byte(v))
		hv := hex.EncodeToString(t.Sum(nil))

		a.Equal(hv, SHA512(v))
	}
}
