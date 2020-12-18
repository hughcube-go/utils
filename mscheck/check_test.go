package mscheck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsTrue(t *testing.T) {
	a := assert.New(t)

	a.True(IsTrue(true))
	a.False(IsTrue(false))

	a.True(IsTrue("on"))
	a.True(IsTrue("ON"))

	a.True(IsTrue("yes"))
	a.True(IsTrue("YES"))

	a.True(IsTrue("true"))
	a.True(IsTrue("TRUE"))

	a.True(IsTrue("1"))

	a.False(IsTrue("1 "))
}

func Test_IsEmail(t *testing.T) {
	a := assert.New(t)

	a.True(IsEmail("123@qq.com"))
	a.False(IsTrue("123qq.com"))
}

func Test_IsUrl(t *testing.T) {
	a := assert.New(t)

	a.True(IsUrl("http://test.com/view/%E9%87%8D%E6%9E%84%E6%9C%8D%E5%8A%A1/job/release_server_v5_console/"))
	a.True(IsUrl("http://192.168.1.10"))
	a.False(IsUrl("192"))
}

func Test_IsNumeric(t *testing.T) {
	a := assert.New(t)

	a.False(IsNumeric("AAA"))
	a.False(IsNumeric("192.168.1.10"))
	a.False(IsNumeric([]interface{}{}))
	a.False(IsNumeric(false))

	a.True(IsNumeric("11111111111111111111111111111"))
	a.True(IsNumeric("11111111111111111111111111111.11111111111111111111111111111"))
	a.True(IsNumeric(0.11111111111111111111111111111))
	a.True(IsNumeric(int64(11111111111111111)))
	a.True(IsNumeric(1))
}

func Test_IsDigit(t *testing.T) {
	a := assert.New(t)

	a.False(IsDigit("AAA"))
	a.False(IsDigit("192.168.1.10"))
	a.False(IsDigit([]interface{}{}))
	a.False(IsDigit(false))

	a.False(IsDigit("11111111111111111111111111111.11111111111111111111111111111"))
	a.False(IsDigit(0.11111111111111111111111111111))

	a.True(IsDigit("11111111111111111"))
	a.True(IsDigit(int64(11111111111111111)))
	a.True(IsDigit(-1))
	a.True(IsDigit("-1"))
}
