package check

import (
	"crypto/tls"
	"crypto/x509/pkix"
	"fmt"
	"net"
	"time"

	"github.com/msfidelis/cassler/src/libs/lookup"
	"github.com/msfidelis/cassler/src/libs/parser"
)

// A Certificate infos to parse
type Certificate struct {
	CommonName            string
	NotBefore             time.Time
	NotAfter              time.Time
	TimeRemain            time.Duration
	IssuingCertificateURL []string
	SignatureAlgorithm    string
	Version               int
	Issuer                pkix.Name
	Subject               pkix.Name
	DNSNames              []string
}

// Cmd to check TLS versions enabled on hosts
func Cmd(url string, port int, dnsServer string) {

	host := parser.ParseHost(url)
	ips := lookup.Lookup(host, dnsServer)

	checkedCertificates := make(map[string]string)
	certificateAuthorities := make(map[string]Certificate)
	certificateList := make(map[string]Certificate)

	fmt.Printf("Checking Certificates: %s on port %d \n", host, port)
	fmt.Printf("\nDNS Lookup on: %s \n\n", dnsServer)

	for _, ip := range ips {

		connConfig := &tls.Config{
			ServerName:         host,
			InsecureSkipVerify: true,
		}

		dialer := net.Dialer{Timeout: 1000000000, Deadline: time.Now().Add(1000000000 + 5*time.Second)}
		connection, err := tls.DialWithDialer(&dialer, "tcp", fmt.Sprintf("[%s]:%d", ip, port), connConfig)

		if err != nil {
			fmt.Printf("%v\n", err)
		} else {

			certificateNegotiationList := connection.ConnectionState().PeerCertificates

			for i := 0; i < len(certificateNegotiationList); i++ {
				cert := certificateNegotiationList[i]

				// Filter Certificate Already Validated
				if _, checked := checkedCertificates[string(cert.Signature)]; checked {
					continue
				}

				checkedCertificates[string(cert.Signature)] = cert.Subject.CommonName

				var certificate Certificate

				certificate.CommonName = cert.Subject.CommonName
				certificate.NotAfter = cert.NotAfter
				certificate.NotBefore = cert.NotBefore
				certificate.TimeRemain = time.Until(cert.NotAfter)
				certificate.SignatureAlgorithm = cert.SignatureAlgorithm.String()
				certificate.IssuingCertificateURL = cert.IssuingCertificateURL
				certificate.Version = cert.Version
				certificate.DNSNames = cert.DNSNames
				certificate.Issuer = cert.Issuer
				certificate.Subject = cert.Subject

				// Filter Certificate Authority
				if cert.IsCA {
					certificateAuthorities[string(cert.Subject.CommonName)] = certificate
					continue
				}

				certificateList[string(cert.Subject.CommonName)] = certificate

			}
		}

	}

	fmt.Printf("Server Certificate: \n")
	for _, data := range certificateList {
		fmt.Printf("Common Name: %s\n", data.CommonName)
		fmt.Printf("Issuer: %s\n", data.Issuer)
		fmt.Printf("Subject: %s\n", data.Subject)
		fmt.Printf("Signature Algorithm: %s\n", data.SignatureAlgorithm)
		fmt.Printf("Created: %s\n", data.NotBefore)
		fmt.Printf("Expires: %s\n", data.NotAfter)
		fmt.Printf("Expiration time: %d days\n", parser.ParseDurationInDays(data.TimeRemain.Hours()))
		fmt.Printf("Certificate Version: %d\n", data.Version)

		if len(data.DNSNames) > 0 {
			fmt.Printf("\nDNS Names: \n")
			for _, dns := range data.DNSNames {
				fmt.Printf("- %s\n", dns)
			}
		}

		if len(data.IssuingCertificateURL) > 0 {
			fmt.Printf("\nIssuing Certificate URL's: \n")
			for _, url := range data.IssuingCertificateURL {
				fmt.Printf("- %s\n", url)
			}
		}
	}

	fmt.Printf("\nServer IP's: \n")
	for _, ip := range ips {
		fmt.Printf("* %s \n", ip)
	}

	fmt.Printf("\nCertificate Authority: \n\n")
	for _, data := range certificateAuthorities {

		fmt.Printf("%s\n", data.CommonName)
		fmt.Printf("Issuer: %s\n", data.Issuer)
		fmt.Printf("Subject: %s\n", data.Subject)
		fmt.Printf("Signature Algorithm: %s\n", data.SignatureAlgorithm)
		fmt.Printf("Created: %s\n", data.NotBefore)
		fmt.Printf("Expires: %s\n", data.NotAfter)
		fmt.Printf("Expiration time: %d days\n", parser.ParseDurationInDays(data.TimeRemain.Hours()))
		fmt.Printf("Certificate Version: %d\n", data.Version)

		if len(data.IssuingCertificateURL) > 0 {
			fmt.Printf("\n\nIssuing Certificate URL's: \n")
			for _, url := range data.IssuingCertificateURL {
				fmt.Printf("- %s\n", url)
			}
		}

		fmt.Printf("\n\n")
	}

}
