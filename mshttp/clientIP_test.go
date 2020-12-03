package mshttp

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_ClientIP(t *testing.T) {
	a := assert.New(t)

	request, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	a.Equal(ClientIP(request), "");

	request.RemoteAddr = "example.com:443"
	a.Equal(ClientIP(request), "");

	request.RemoteAddr = "192.168.1.1:443"
	a.Equal(ClientIP(request), "192.168.1.1");

	request.Header.Set("X-Real-Ip", "192.168.1.1")
	a.Equal(ClientIP(request), "192.168.1.1");

	request.Header.Set("X-Forwarded-For", "192.168.1.1")
	a.Equal(ClientIP(request), "192.168.1.1");

	request.Header.Set("X-Forwarded-For", "192.168.1.1,192.168.1.2")
	a.Equal(ClientIP(request), "192.168.1.1");
}
