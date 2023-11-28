package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/doktorupnos/crow/backend/internal/app"
)

func TestHealthCheck(t *testing.T) {
	t.Run("returns 200", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(app.HealthCheck))
		defer server.Close()

		resp, err := http.Get(server.URL)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		got := resp.StatusCode
		want := http.StatusOK
		if got != want {
			t.Errorf("\ngot: %s\nwant: %s\n", http.StatusText(got), http.StatusText(want))
		}
	})
}
