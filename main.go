package main

import (
	"flag"
	"os"
	"strings"

	"github.com/msfidelis/cassler/cmd/check"
	"github.com/msfidelis/cassler/cmd/tls"
)

func main() {
	url := flag.String("url", "", "URL to validate SSL certificate,ex: https://google.com")
	port := flag.Int("port", 443, "Server port, default: 443")
	dnsServer := flag.String("dns", "8.8.8.8", "DNS Server, default 8.8.8.8")
	reverseLookup := flag.Bool("lookup", false, "Check reverse DNS Lookup for hosts IP's")
	mode := flag.String("mode", "check", "Actions; Default: `check`; Available options `check` for check certificates, `tls` to test TLS connection, `scan` for complete checks on hosts")
	flag.Parse()

	switch strings.ToLower(*mode) {
	case "check":
		if *url == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
		check.Cmd(*url, *port, *dnsServer)
	case "tls":
		if *url == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
		tls.Cmd(*url, *port, *dnsServer, *reverseLookup)
	case "scan":
		if *url == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
		check.Cmd(*url, *port, *dnsServer)
		tls.Cmd(*url, *port, *dnsServer, true)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
