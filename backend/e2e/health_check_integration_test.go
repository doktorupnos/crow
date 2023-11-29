package integration_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/doktorupnos/crow/backend/internal/app"
)

func TestHealthCheckIntegration(t *testing.T) {
	server := NewTestServer(app.HealthCheck)
	defer server.Close()

	client := http.Client{Timeout: time.Minute}

	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer resp.Body.Close()

	got := resp.StatusCode
	want := http.StatusOK

	assertEqual(t, got, want)
}
