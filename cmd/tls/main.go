package tls

import (
	"crypto/tls"
	"fmt"

	"github.com/msfidelis/cassler/src/libs/lookup"
	"github.com/msfidelis/cassler/src/libs/parser"
	"github.com/msfidelis/cassler/src/libs/tls_check"
)

type Validation struct {
	Ip    string
	Host  string
	TLS10 bool
	TLS11 bool
	TLS12 bool
	TLS13 bool
}

func Cmd(url string, port int, dns_server string) {
	host := parser.ParseHost(url)
	ips := lookup.Lookup(host, dns_server)

	tls_versions := map[string]uint16{
		"tls1.0": tls.VersionTLS10,
		"tls1.1": tls.VersionTLS11,
		"tls1.2": tls.VersionTLS12,
		"tls1.3": tls.VersionTLS13,
	}

	validation_list := make(map[string]Validation)

	fmt.Printf("\nTesting TLS Versions: %s on port %d \n", host, port)
	fmt.Printf("\nDNS Lookup on: %s \n\n", dns_server)

	for _, ip := range ips {

		var validation Validation

		validation.Ip = ip
		validation.Host = host
		validation.TLS10 = tls_check.Check(host, ip, port, tls_versions["tls1.0"])
		validation.TLS11 = tls_check.Check(host, ip, port, tls_versions["tls1.1"])
		validation.TLS12 = tls_check.Check(host, ip, port, tls_versions["tls1.2"])
		validation.TLS13 = tls_check.Check(host, ip, port, tls_versions["tls1.3"])

		validation_list[fmt.Sprintf("%v", ip)] = validation
	}

	for ip, validation := range validation_list {
		fmt.Printf("TLS Versions Enabled on %v: \n", ip)
		fmt.Printf("- tls1.0: %v \n", validation.TLS10)
		fmt.Printf("- tls1.1: %v \n", validation.TLS11)
		fmt.Printf("- tls1.2: %v \n", validation.TLS12)
		fmt.Printf("- tls1.3: %v \n", validation.TLS13)
		fmt.Printf("\n")
	}

}
