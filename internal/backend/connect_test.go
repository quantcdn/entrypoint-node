package backend_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/quantcdn/entrypoint-node/internal/backend"
	"github.com/stretchr/testify/assert"
)

func TestConnect_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	url := server.URL
	delay := 1
	retry := 3

	result := backend.Connect(url, delay, retry)
	assert.True(t, result)
}

func TestConnect_Retries(t *testing.T) {
	// Create a test HTTP server that initially returns an error,
	// then starts returning OK after a few retries
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts <= 3 {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer server.Close()

	url := server.URL
	delay := 1
	retry := 5 // eSt a number of retries

	result := backend.Connect(url, delay, retry)
	assert.True(t, result)
}

func TestConnect_Fail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()
	url := server.URL
	delay := 1
	retry := 3

	result := backend.Connect(url, delay, retry)
	assert.False(t, result)
}

func TestConnect_InvalidURL(t *testing.T) {
	url := "invalid-url"
	delay := 1
	retry := 0

	result := backend.Connect(url, delay, retry)
	assert.False(t, result)
}

func TestConnect_NoRetry(t *testing.T) {
	// Create a test HTTP server that always returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	url := server.URL

	delay := 1
	retry := 0 // No retries

	result := backend.Connect(url, delay, retry)
	assert.False(t, result)
}

func TestConnect_Envar(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.Path, "jsonapi"))
		w.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	url := server.URL
	os.Setenv("NEXT_PUBLIC_DRUPAL_BASE_URL", url)
	result := backend.Connect(url, 1, 3)

	assert.True(t, result)
}
