package main

import (
	"flag"
	"os"
	"strings"

	"github.com/msfidelis/cassler/cmd/check"
)

func main() {
	url := flag.String("url", "", "URL to validate certificate,ex: https://google.com")
	port := flag.Int("port", 443, "Server port, default: 443")
	mode := flag.String("mode", "check", "Actions; Default: `check`; Available options `check` for check certificates, `tls` to test TLS connection")
	flag.Parse()

	switch strings.ToLower(*mode) {
	case "check":
		if *url == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
		check.Cmd(*url, *port)
		break
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
