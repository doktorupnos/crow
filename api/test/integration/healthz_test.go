package integration

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/doktorupnos/crow/api/internal/app"
)

func TestHealthz(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(app.Healthz))
	defer server.Close()

	client := server.Client()
	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatal("Get() unexpected error:", err)
	}
	defer resp.Body.Close()

	gotStatus := resp.StatusCode
	wantStatus := http.StatusOK
	if gotStatus != wantStatus {
		t.Errorf("Healthz() got status code %d, want %d", gotStatus, wantStatus)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("ReadAll(resp.Body) unexpected error:", err)
	}

	gotBody := string(data)
	wantBody := "OK"
	if gotBody != wantBody {
		t.Errorf("Healthz() got body %q, want %q", gotBody, wantBody)
	}
}
