package tlscheck

import (
	"crypto/tls"
	"testing"

	"github.com/msfidelis/cassler/src/libs/lookup"
)

func TestTLS10Enabled(t *testing.T) {
	ips := lookup.Lookup("tls-v1-0.badssl.com", "8.8.8.8")
	enabled := Check("tls-v1-0.badssl.com", ips[0], 1010, tls.VersionTLS10)

	if enabled != true {
		t.Errorf("Expected true, got %v", enabled)
	}
}

func TestTLS11EnabledOn10(t *testing.T) {
	ips := lookup.Lookup("tls-v1-0.badssl.com", "8.8.8.8")
	enabled := Check("tls-v1-0.badssl.com", ips[0], 1010, tls.VersionTLS11)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS12EnabledOn10(t *testing.T) {
	ips := lookup.Lookup("tls-v1-0.badssl.com", "8.8.8.8")
	enabled := Check("tls-v1-0.badssl.com", ips[0], 1010, tls.VersionTLS12)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS13EnabledOn10(t *testing.T) {
	ips := lookup.Lookup("tls-v1-0.badssl.com", "8.8.8.8")
	enabled := Check("tls-v1-0.badssl.com", ips[0], 1010, tls.VersionTLS13)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS11Enabled(t *testing.T) {
	ips := lookup.Lookup("tls-v1-1.badssl.com", "8.8.8.8")
	enabled := Check("tls-v1-1.badssl.com", ips[0], 1011, tls.VersionTLS11)

	if enabled != true {
		t.Errorf("Expected true, got %v", enabled)
	}
}

func TestTLS10EnabledOn11(t *testing.T) {
	ips := lookup.Lookup("tls-v1-1.badssl.com", "8.8.8.8")
	enabled := Check("tls-v1-1.badssl.com", ips[0], 1011, tls.VersionTLS10)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS12EnabledOn11(t *testing.T) {
	ips := lookup.Lookup("tls-v1-1.badssl.com", "8.8.8.8")
	enabled := Check("tls-v1-1.badssl.com", ips[0], 1011, tls.VersionTLS12)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS13EnabledOn11(t *testing.T) {
	ips := lookup.Lookup("tls-v1-1.badssl.com", "8.8.8.8")
	enabled := Check("tls-v1-1.badssl.com", ips[0], 1011, tls.VersionTLS13)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS12Enabled(t *testing.T) {
	ips := lookup.Lookup("tls-v1-2.badssl.com", "8.8.8.8")
	enabled := Check("tls-v1-2.badssl.com", ips[0], 1012, tls.VersionTLS12)

	if enabled != true {
		t.Errorf("Expected true, got %v", enabled)
	}
}

func TestTLS10EnabledOn12(t *testing.T) {
	ips := lookup.Lookup("tls-v1-2.badssl.com", "8.8.8.8")
	enabled := Check("tls-v1-2.badssl.com", ips[0], 1012, tls.VersionTLS10)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS11EnabledOn12(t *testing.T) {
	ips := lookup.Lookup("tls-v1-2.badssl.com", "8.8.8.8")
	enabled := Check("tls-v1-2.badssl.com", ips[0], 1012, tls.VersionTLS11)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS13EnabledOn12(t *testing.T) {
	ips := lookup.Lookup("tls-v1-2.badssl.com", "8.8.8.8")
	enabled := Check("tls-v1-2.badssl.com", ips[0], 1012, tls.VersionTLS13)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}
