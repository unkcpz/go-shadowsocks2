package main

import (
	"testing"
)

func TestParseURL(t *testing.T) {
	tests := []struct {
		s            string
		wantAddr     string
		wantCipher   string
		wantPassword string
		wantErr      error
	}{
		{
			"ss://AEAD_CHACHA20_POLY1305:passwd@[0.0.0.0]:8000",
			"[0.0.0.0]:8000",
			"AEAD_CHACHA20_POLY1305",
			"passwd",
			nil,
		},
	}

	for _, test := range tests {
		gotAddr, gotCipher, gotPassword, gotErr := parseURL(test.s)
		if gotAddr != test.wantAddr {
			t.Errorf("parseURL(%q) got Addr %s, want %s",
				test.s, gotAddr, test.wantAddr)
		}
		if gotCipher != test.wantCipher {
			t.Errorf("parseURL(%q) got Cipher %s, want %s",
				test.s, gotCipher, test.wantCipher)
		}
		if gotPassword != test.wantPassword {
			t.Errorf("parseURL(%q) got Password %s, want %s",
				test.s, gotPassword, test.wantPassword)
		}
		if gotErr != test.wantErr {
			t.Errorf("parseURL(%q) return %v, want %v",
				test.s, gotErr, test.wantErr)
		}
	}
}
