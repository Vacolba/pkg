package web

import (
	"net"
	"net/http"
)

// GetHTTPRequestAddress returns client IP from an http.Request
func GetHTTPRequestAddress(r *http.Request) string {
	address := r.Header.Get("X-Forwarded-For")

	if len(address) == 0 {
		address = r.Header.Get("X-Real-IP")
	}

	if len(address) == 0 {
		address, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	return address
}
