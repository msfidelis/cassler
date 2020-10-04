package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type Certificate struct {
	CommonName         string
	NotBefore          time.Time
	NotAfter           time.Time
	TimeRemain         time.Duration
	SignatureAlgorithm string
}

func main() {

	url := flag.String("url", "", "URL to validate certificate,ex: https://google.com")
	port := flag.Int("port", 443, "Server port, default: 443")
	flag.Parse()

	if *url == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	host := ParseHost(*url)
	ips := Lookup(host)

	checked_certificates := make(map[string]string)
	certificate_authorities := make(map[string]string)

	certificate_list := make(map[string]Certificate)

	for _, ip := range ips {

		conn_config := &tls.Config{
			ServerName:         host,
			InsecureSkipVerify: true,
		}

		dialer := net.Dialer{Timeout: 1000000000, Deadline: time.Now().Add(1000000000 + 5*time.Second)}
		connection, err := tls.DialWithDialer(&dialer, "tcp", fmt.Sprintf("[%s]:%d", ip, *port), conn_config)
		defer connection.Close()

		if err != nil {
			fmt.Printf("%v\n", err)
		} else {

			certificate_negotiation_list := connection.ConnectionState().PeerCertificates

			for i := 0; i < len(certificate_negotiation_list); i++ {
				cert := certificate_negotiation_list[i]

				// Filter Certificate Already Validated
				if _, checked := checked_certificates[string(cert.Signature)]; checked {
					continue
				}
				checked_certificates[string(cert.Signature)] = cert.Subject.CommonName

				// Filter Certificate Authority
				if cert.IsCA {
					certificate_authorities[string(cert.Subject.CommonName)] = cert.Subject.CommonName
					continue
				}

				var certificate Certificate

				certificate.CommonName = cert.Subject.CommonName
				certificate.NotAfter = cert.NotAfter
				certificate.NotBefore = cert.NotBefore
				certificate.TimeRemain = cert.NotAfter.Sub(time.Now())
				certificate.SignatureAlgorithm = cert.SignatureAlgorithm.String()

				certificate_list[string(cert.Subject.CommonName)] = certificate

			}
		}

	}

	fmt.Printf("Resolving: %s on port %d \n\n", host, *port)

	fmt.Printf("Server Certificate: \n")
	for _, data := range certificate_list {
		fmt.Printf("Common Name: %s\n", data.CommonName)
		fmt.Printf("Signature Algorithm: %s\n", data.SignatureAlgorithm)
		fmt.Printf("Created: %s\n", data.NotBefore)
		fmt.Printf("Expires: %s\n", data.NotAfter)
		fmt.Printf("Expiration time: %d days\n", ParseDurationInDays(data.TimeRemain.Hours()))
	}

	fmt.Printf("\nServer IP's: \n")
	for _, ip := range ips {
		fmt.Printf("* %s \n", ip)
	}

	fmt.Printf("\nCertificate Authority: \n")
	for _, ca := range certificate_authorities {
		fmt.Printf("* %s \n", ca)
	}

}

func Lookup(url string) []net.IP {
	timer := time.NewTimer(1000000000)
	ch := make(chan []net.IP, 1)
	go func() {
		r, err := net.LookupIP(url)
		if err != nil {
			fmt.Printf("%v", err)
		}
		ch <- r
	}()
	select {
	case ips := <-ch:
		return ips
	case <-timer.C:
		fmt.Printf("timeout resolving %s\n", url)
	}
	return make([]net.IP, 0)
}

func ParseHost(url string) string {
	var result string
	result = strings.ToLower(url)
	result = strings.TrimPrefix(result, "https://")
	result = strings.TrimPrefix(result, "http://")
	result = strings.TrimPrefix(result, "ftp://")
	result = strings.TrimPrefix(result, "ws://")
	return result
}

func ParseDurationInDays(duration float64) int {
	floatDays := duration / 24
	parsedDays := int(floatDays)
	return parsedDays
}
