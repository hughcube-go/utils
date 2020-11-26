package msip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsIPv4(t *testing.T) {
	a := assert.New(t)

	a.True(IsIPv4("192.168.1.10"))
	a.False(IsIPv4("192.168.1."))
}

func Test_IsIPv6(t *testing.T) {
	a := assert.New(t)

	a.True(IsIPv6("2409:8954:3054:58ff:c009:c9e1:3af5:c13e"))
	a.False(IsIPv6("2409:8954:3054:58ff:c009:c9e1:3af5:"))
}

func Test_IsIP(t *testing.T) {
	a := assert.New(t)

	a.True(IsIP("2409:8954:3054:58ff:c009:c9e1:3af5:c13e"))
	a.True(IsIP("192.168.1.10"))
}

func Test_InRange(t *testing.T) {
	a := assert.New(t)

	a.False(InRange("2409:8954:3054:58ff:c009:c9e1:3af5:c13e", "192.168.1.10"))
	a.False(InRange("192.168.1.21/24", "192.168.1.21/32"))

	a.True(InRange("192.168.1.0", "192.168.0.0/22"))
	a.True(InRange("192.168.1.21", "192.168.1.0/24"))
	a.True(InRange("192.168.1.21", "192.168.1.0/24"))
	a.True(InRange("192.168.1.21", "192.168.1.21"))
}
