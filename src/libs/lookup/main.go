package lookup

import (
	"fmt"
	"net"
	"time"
)

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
