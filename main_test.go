package main_test

import (
	"testing"

	"github.com/msfidelis/cassler/src/libs/lookup"
	"github.com/msfidelis/cassler/src/libs/parser"
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
