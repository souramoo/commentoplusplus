package main

import (
	"os"
	"strings"
	"testing"
)

func TestEmailStripBasics(t *testing.T) {
	tests := map[string]string{
		"test@example.com":              "test@example.com",
		"test+strip@example.com":        "test@example.com",
		"test+strip+strip2@example.com": "test@example.com",
	}

	for in, out := range tests {
		if emailStrip(in) != out {
			t.Errorf("for in=%s expected out=%s got out=%s", in, out, emailStrip(in))
			return
		}
	}
}

func TestAddHttpIfAbsent(t *testing.T) {
	tests := map[string]string{
		"http://example.com": "http://example.com",
		"https://example.com": "https://example.com",
		"example.com":         "http://example.com",
		"example.com/":        "https://example.com/",
	}

	for in, out := range tests {
		if strings.HasSuffix(in, "/") {
			os.Setenv("SSL", "true")
		} else {
			os.Setenv("SSL", "false")
		}
		if addHttpIfAbsent(in) != out {
			t.Errorf("for in=%s expected out=%s got out=%s", in, out, addHttpIfAbsent(in))
			return
		}
	}
}