package parser

import (
	"testing"
)

const expected_url = "google.com"

func TestParserNormalUrlsFTP(t *testing.T) {
	ftp := ParseHost("ftp://google.com")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsFTP(t *testing.T) {
	ftp := ParseHost("FtP://google.Com")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsFTPUpper(t *testing.T) {
	ftp := ParseHost("FTP://GOOGLE.COM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserNormalUrlsHTTP(t *testing.T) {
	ftp := ParseHost("http://google.com")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsHTTPUpper(t *testing.T) {
	ftp := ParseHost("HTTP://GOOGLE.COM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsHTTPCaMeL(t *testing.T) {
	ftp := ParseHost("HTtP://gOoGlE.CoM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserNormalUrlsHTTPS(t *testing.T) {
	ftp := ParseHost("https://google.com")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsHTTPSUpper(t *testing.T) {
	ftp := ParseHost("HTTPS://GOOGLE.COM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsHTTPSCaMeL(t *testing.T) {
	ftp := ParseHost("HTtPS://gOoGlE.CoM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserNormalUrlsWS(t *testing.T) {
	ftp := ParseHost("ws://google.com")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsWSUpper(t *testing.T) {
	ftp := ParseHost("WS://GOOGLE.COM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserInvalidUrlsWSCaMeL(t *testing.T) {
	ftp := ParseHost("Ws://gOoGlE.CoM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserNormalUrls(t *testing.T) {
	ftp := ParseHost("google.com")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserNormalUPPER(t *testing.T) {
	ftp := ParseHost("GOOGLE.COM")

	if ftp != expected_url {
		t.Errorf("Expected %s, got %s", expected_url, ftp)
	}
}

func TestParserHours(t *testing.T) {
	var hours = 13049.651703628333
	var expected_days = 543
	days := ParseDurationInDays(hours)

	if days != expected_days {
		t.Errorf("Expected %d, got %d", expected_days, days)
	}
}
