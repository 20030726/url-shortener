package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShortenHandlerUsesRequestHost(t *testing.T) {
	// Reset url store between tests
	mu.Lock()
	urlStore = make(map[string]string)
	mu.Unlock()

	body := bytes.NewBufferString(`{"url":"https://example.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/shorten", body)
	req.Host = "myhost.example.com:9090"
	w := httptest.NewRecorder()

	shortenHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp map[string]string
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	shortURL, ok := resp["short_url"]
	if !ok {
		t.Fatal("response missing short_url field")
	}

	if !strings.HasPrefix(shortURL, "http://myhost.example.com:9090/") {
		t.Errorf("expected short_url to start with http://myhost.example.com:9090/, got %s", shortURL)
	}
}

func TestShortenHandlerUsesHTTPSWhenTLS(t *testing.T) {
	// Reset url store between tests
	mu.Lock()
	urlStore = make(map[string]string)
	mu.Unlock()

	body := bytes.NewBufferString(`{"url":"https://example.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/shorten", body)
	req.Host = "secure.example.com"
	// Simulate a TLS connection by setting a non-nil TLS field
	req.TLS = &tls.ConnectionState{}
	w := httptest.NewRecorder()

	shortenHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp map[string]string
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	shortURL, ok := resp["short_url"]
	if !ok {
		t.Fatal("response missing short_url field")
	}

	if !strings.HasPrefix(shortURL, "https://secure.example.com/") {
		t.Errorf("expected short_url to start with https://secure.example.com/, got %s", shortURL)
	}
}
