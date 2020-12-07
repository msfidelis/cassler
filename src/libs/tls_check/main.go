package tls_check

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

func Check(host string, ip string, port int, tls_version uint16) bool {
	conn_config := &tls.Config{
		ServerName:         host,
		InsecureSkipVerify: false,
		MinVersion:         tls_version,
		MaxVersion:         tls_version,
	}
	dialer := net.Dialer{Timeout: 1000000000, Deadline: time.Now().Add(1000000000 + 5*time.Second)}
	_, err := tls.DialWithDialer(&dialer, "tcp", fmt.Sprintf("[%s]:%d", ip, port), conn_config)

	if err != nil {
		return false
	} else {
		return true
	}
}
