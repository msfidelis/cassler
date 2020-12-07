package main

import (
	"flag"
	"os"
	"strings"

	"github.com/msfidelis/cassler/cmd/check"
	"github.com/msfidelis/cassler/cmd/tls"
)

func main() {
	url := flag.String("url", "", "URL to validate certificate,ex: https://google.com")
	port := flag.Int("port", 443, "Server port, default: 443")
	dns_server := flag.String("dns", "8.8.8.8", "DNS Server, default 8.8.8.8")
	mode := flag.String("mode", "check", "Actions; Default: `check`; Available options `check` for check certificates, `tls` to test TLS connection, `scan` for complete checks on hosts")
	flag.Parse()

	switch strings.ToLower(*mode) {
	case "check":
		if *url == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
		check.Cmd(*url, *port, *dns_server)
		break
	case "tls":
		if *url == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
		tls.Cmd(*url, *port, *dns_server)
		break
	case "scan":
		if *url == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
		check.Cmd(*url, *port, *dns_server)
		tls.Cmd(*url, *port, *dns_server)
		break
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
