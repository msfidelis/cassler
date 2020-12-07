package lookup

import (
	"context"
	"fmt"
	"net"
	"time"
)

func Lookup(url string, dns_server string) []string {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, "udp", fmt.Sprintf("%s:%s", dns_server, "53"))
		},
	}
	ip, _ := r.LookupHost(context.Background(), url)

	return ip
}
