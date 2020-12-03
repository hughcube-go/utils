package mshttp

import (
	"net"
	"net/http"
	"strings"
)

func ClientIP(r *http.Request) string {
	clientIP := r.Header.Get("X-Forwarded-For")
	if items := strings.Split(clientIP, ","); 0 < len(items) {
		clientIP = strings.TrimSpace(items[0])
		if clientIP != "" {
			return clientIP
		}
	}

	clientIP = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if clientIP != "" {
		return clientIP
	}

	if clientIP, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		if ip := net.ParseIP(clientIP); ip != nil {
			return clientIP
		}
	}

	return ""
}
