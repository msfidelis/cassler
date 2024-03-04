package tls

import (
	"crypto/tls"
	"fmt"

	"github.com/msfidelis/cassler/src/libs/lookup"
	"github.com/msfidelis/cassler/src/libs/parser"
	"github.com/msfidelis/cassler/src/libs/tlscheck"
)

// A Validation types to print
type Validation struct {
	IP    string
	Host  string
	TLS10 bool
	TLS11 bool
	TLS12 bool
	TLS13 bool
}

// Cmd to check TLS versions enabled on hosts
func Cmd(url string, port int, dnsServer string, reverseLookup bool) {
	host := parser.ParseHost(url)
	ips := lookup.Lookup(host, dnsServer)

	tlsVersions := map[string]uint16{
		"tls1.0": tls.VersionTLS10,
		"tls1.1": tls.VersionTLS11,
		"tls1.2": tls.VersionTLS12,
		"tls1.3": tls.VersionTLS13,
	}

	validationList := make(map[string]Validation)

	fmt.Printf("\nTesting TLS Versions: %s on port %d \n", host, port)
	fmt.Printf("\nDNS Lookup on: %s \n\n", dnsServer)

	for _, ip := range ips {

		var validation Validation

		validation.IP = ip
		validation.Host = host
		validation.TLS10 = tlscheck.Check(host, ip, port, tlsVersions["tls1.0"])
		validation.TLS11 = tlscheck.Check(host, ip, port, tlsVersions["tls1.1"])
		validation.TLS12 = tlscheck.Check(host, ip, port, tlsVersions["tls1.2"])
		validation.TLS13 = tlscheck.Check(host, ip, port, tlsVersions["tls1.3"])

		validationList[fmt.Sprintf("%v", ip)] = validation
	}

	for _, validation := range validationList {
		fmt.Printf("TLS Versions Enabled on %v: \n", validation.IP)
		fmt.Printf("- tls1.0: %v \n", validation.TLS10)
		fmt.Printf("- tls1.1: %v \n", validation.TLS11)
		fmt.Printf("- tls1.2: %v \n", validation.TLS12)
		fmt.Printf("- tls1.3: %v \n", validation.TLS13)
		fmt.Printf("\n")
	}

	if reverseLookup {

		for _, ip := range ips {
			fmt.Printf("\nStarting reverse DNS Lookup on:  %v: \n", ip)
			domains, err := lookup.ReverseLookup(ip)
			if err != nil {
				fmt.Printf("\nError to reverse lookup on %v\n", ip)
				fmt.Println(err)
			} else {
				for _, ad := range domains {
					fmt.Printf("%v:  %v\n", ip, ad)
				}
			}
		}

		// for _, ip := range ips {
		// 	fmt.Printf("\nStarting reverse DNS Lookup on:  %v: \n", ip)
		// 	addr, err := net.LookupAddr(ip)

		// 	if err != nil {
		// 		fmt.Printf("\nError to reverse lookup on %v\n", ip)
		// 		fmt.Println(err)
		// 	} else {
		// 		for _, ad := range addr {
		// 			fmt.Printf("%v:  %v\n", ip, ad)
		// 		}
		// 	}
		// }
	}

}
