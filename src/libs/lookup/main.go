package lookup

import (
	"context"
	"fmt"
	"net"
	"time"
)

func Lookup(url string, dns_server string) []string {
	timer := time.NewTimer(1000000000)
	ch := make(chan []string, 1)
	go func() {
		r := &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: time.Millisecond * time.Duration(10000),
				}
				return d.DialContext(ctx, "udp", fmt.Sprintf("%s:%s", dns_server, "53"))
			},
		}
		ip, err := r.LookupHost(context.Background(), url)
		if err != nil {
			fmt.Printf("%v", err)
		}
		ch <- ip
	}()
	select {
	case ips := <-ch:
		return ips
	case <-timer.C:
		fmt.Printf("timeout resolving %s\n", url)
	}
	return make([]string, 0)
}
