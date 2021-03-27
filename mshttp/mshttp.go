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

func IsSecureConnection(request *http.Request) bool {
	headers := map[string][]string{}
	headers["X-Forwarded-Proto"] = []string{"https"} // Common
	headers["Front-End-Https"] = []string{"on"}      // Microsoft

	for name, values := range headers {
		for _, requestValue := range request.Header.Values(name) {
			for _, value := range values {
				if strings.EqualFold(value, requestValue) {
					return true
				}
			}
		}
	}

	return false
}
