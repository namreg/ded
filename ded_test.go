package ded

import (
	"testing"
)

func TestIsDisposableEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		want    bool
		wantErr error
	}{
		{"not disposable email", "eml@google.com", false, nil},
		{"disposable email", "eml@mail.wtf", true, nil},
		{"ignore spaces", "   eml@mail.wtf   ", true, nil},
		{"case insensitive", "eml@MAIL.wtf", true, nil},
		{"malformed email", "eml.com", false, ErrEmailMalformed},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsDisposableEmail(tt.email)
			if err != tt.wantErr {
				t.Fatalf("IsDisposableEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("IsDisposableEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDisposableDomain(t *testing.T) {
	tests := []struct {
		name    string
		domain  string
		want    bool
		wantErr error
	}{
		{"not disposable domain", "google.com", false, nil},
		{"disposable domain", "mail.wtf", true, nil},
		{"ignore spaces", "   mail.wtf   ", true, nil},
		{"case insensitive", "MAIL.wtf", true, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsDisposableDomain(tt.domain)
			if err != tt.wantErr {
				t.Fatalf("TestIsDisposableDomain() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("TestIsDisposableDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
