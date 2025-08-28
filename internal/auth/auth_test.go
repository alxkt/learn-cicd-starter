package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_Simple(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey test-key")
	got, err := GetAPIKey(h)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if got != "test-key" {
		t.Errorf("expected key: %q, got: %q", "test-key", got)
	}
}
