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
