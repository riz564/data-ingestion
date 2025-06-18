package data_collection

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Test API timeout
func TestGetPosts_Timeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()

	timeout := 200 * time.Millisecond
	_, err := GetPosts(ts.URL, timeout)
	if err == nil {
		t.Error("expected timeout error, got nil")
	}
}

// Test non-200 response
func TestGetPosts_Non200(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	_, err := GetPosts(ts.URL, 1*time.Second)
	if err == nil || err.Error() != "non-200 response" {
		t.Errorf("expected non-200 response error, got %v", err)
	}
}

// Test invalid JSON response
func TestGetPosts_InvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`not json`))
	}))
	defer ts.Close()

	_, err := GetPosts(ts.URL, 1*time.Second)
	if err == nil {
		t.Error("expected JSON decode error, got nil")
	}
}
