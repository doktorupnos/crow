package integration_test

import (
	"net/http"
	"testing"
)

func TestHealthCheckIntegration(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	resp, err := client.Get(server.URL + "/healthz")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer resp.Body.Close()

	got := resp.StatusCode
	want := http.StatusOK

	assertEqual(t, got, want)
}
