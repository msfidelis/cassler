package tlscheck

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

// Check if TLS version is enabled on host
func Check(host string, ip string, port int, tlsVersion uint16) bool {
	connConfig := &tls.Config{
		ServerName:         host,
		InsecureSkipVerify: false,
		MinVersion:         tlsVersion,
		MaxVersion:         tlsVersion,
	}
	dialer := net.Dialer{Timeout: 1000000000, Deadline: time.Now().Add(1000000000 + 5*time.Second)}
	_, err := tls.DialWithDialer(&dialer, "tcp", fmt.Sprintf("[%s]:%d", ip, port), connConfig)

	return err == nil
}
