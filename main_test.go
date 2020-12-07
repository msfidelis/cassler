package main_test

import (
	"crypto/tls"
	"testing"

	"github.com/msfidelis/cassler/src/libs/lookup"
	"github.com/msfidelis/cassler/src/libs/parser"
	"github.com/msfidelis/cassler/src/libs/tls_check"
)

const expected_url = "google.com"

func TestParserNormalUrlsFTP(t *testing.T) {
	ftp := parser.ParseHost("ftp://google.com")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsFTP(t *testing.T) {
	ftp := parser.ParseHost("FtP://google.Com")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsFTPUpper(t *testing.T) {
	ftp := parser.ParseHost("FTP://GOOGLE.COM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserNormalUrlsHTTP(t *testing.T) {
	ftp := parser.ParseHost("http://google.com")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsHTTPUpper(t *testing.T) {
	ftp := parser.ParseHost("HTTP://GOOGLE.COM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsHTTPCaMeL(t *testing.T) {
	ftp := parser.ParseHost("HTtP://gOoGlE.CoM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserNormalUrlsHTTPS(t *testing.T) {
	ftp := parser.ParseHost("https://google.com")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsHTTPSUpper(t *testing.T) {
	ftp := parser.ParseHost("HTTPS://GOOGLE.COM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsHTTPSCaMeL(t *testing.T) {
	ftp := parser.ParseHost("HTtPS://gOoGlE.CoM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserNormalUrlsWS(t *testing.T) {
	ftp := parser.ParseHost("ws://google.com")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsWSUpper(t *testing.T) {
	ftp := parser.ParseHost("WS://GOOGLE.COM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsWSCaMeL(t *testing.T) {
	ftp := parser.ParseHost("Ws://gOoGlE.CoM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserNormalUrls(t *testing.T) {
	ftp := parser.ParseHost("google.com")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserNormalUPPER(t *testing.T) {
	ftp := parser.ParseHost("GOOGLE.COM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserHours(t *testing.T) {
	var hours = 13049.651703628333
	var expected_days = 543
	days := parser.ParseDurationInDays(hours)

	if days != expected_days {
		t.Errorf("Expected %d, got %d", expected_days, days)
	}
}

// Lookup
func TestLookupIsResolving(t *testing.T) {
	ips := lookup.Lookup(expected_url, "8.8.8.8")

	if len(ips) == 0 {
		t.Errorf("Expected len > 0, got %d", len(ips))
	}
}

//TLS Check
func TestTLS10Enabled(t *testing.T) {
	ips := lookup.Lookup("tls-v1-0.badssl.com", "8.8.8.8")
	enabled := tls_check.Check("tls-v1-0.badssl.com", ips[0], 1010, tls.VersionTLS10)

	if enabled != true {
		t.Errorf("Expected true, got %v", enabled)
	}
}

func TestTLS11EnabledOn10(t *testing.T) {
	ips := lookup.Lookup("tls-v1-0.badssl.com", "8.8.8.8")
	enabled := tls_check.Check("tls-v1-0.badssl.com", ips[0], 1010, tls.VersionTLS11)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS12EnabledOn10(t *testing.T) {
	ips := lookup.Lookup("tls-v1-0.badssl.com", "8.8.8.8")
	enabled := tls_check.Check("tls-v1-0.badssl.com", ips[0], 1010, tls.VersionTLS12)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS13EnabledOn10(t *testing.T) {
	ips := lookup.Lookup("tls-v1-0.badssl.com", "8.8.8.8")
	enabled := tls_check.Check("tls-v1-0.badssl.com", ips[0], 1010, tls.VersionTLS13)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS11Enabled(t *testing.T) {
	ips := lookup.Lookup("tls-v1-1.badssl.com", "8.8.8.8")
	enabled := tls_check.Check("tls-v1-1.badssl.com", ips[0], 1011, tls.VersionTLS11)

	if enabled != true {
		t.Errorf("Expected true, got %v", enabled)
	}
}

func TestTLS10EnabledOn11(t *testing.T) {
	ips := lookup.Lookup("tls-v1-1.badssl.com", "8.8.8.8")
	enabled := tls_check.Check("tls-v1-1.badssl.com", ips[0], 1011, tls.VersionTLS10)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS12EnabledOn11(t *testing.T) {
	ips := lookup.Lookup("tls-v1-1.badssl.com", "8.8.8.8")
	enabled := tls_check.Check("tls-v1-1.badssl.com", ips[0], 1011, tls.VersionTLS12)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS13EnabledOn11(t *testing.T) {
	ips := lookup.Lookup("tls-v1-1.badssl.com", "8.8.8.8")
	enabled := tls_check.Check("tls-v1-1.badssl.com", ips[0], 1011, tls.VersionTLS13)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS12Enabled(t *testing.T) {
	ips := lookup.Lookup("tls-v1-2.badssl.com", "8.8.8.8")
	enabled := tls_check.Check("tls-v1-2.badssl.com", ips[0], 1012, tls.VersionTLS12)

	if enabled != true {
		t.Errorf("Expected true, got %v", enabled)
	}
}

func TestTLS10EnabledOn12(t *testing.T) {
	ips := lookup.Lookup("tls-v1-2.badssl.com", "8.8.8.8")
	enabled := tls_check.Check("tls-v1-2.badssl.com", ips[0], 1012, tls.VersionTLS10)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS11EnabledOn12(t *testing.T) {
	ips := lookup.Lookup("tls-v1-2.badssl.com", "8.8.8.8")
	enabled := tls_check.Check("tls-v1-2.badssl.com", ips[0], 1012, tls.VersionTLS11)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}

func TestTLS13EnabledOn12(t *testing.T) {
	ips := lookup.Lookup("tls-v1-2.badssl.com", "8.8.8.8")
	enabled := tls_check.Check("tls-v1-2.badssl.com", ips[0], 1012, tls.VersionTLS13)

	if enabled == true {
		t.Errorf("Expected false, got %v", enabled)
	}
}
