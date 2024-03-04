package lookup

import (
	"testing"
)

func TestLookup(t *testing.T) {
	dnsServer := "8.8.8.8" // Google Public DNS para o teste
	url := "google.com"
	ips := Lookup(url, dnsServer)
	if len(ips) == 0 {
		t.Errorf("Lookup não retornou nenhum IP para %s", url)
	}
}
