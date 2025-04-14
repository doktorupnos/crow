package integration_test

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/doktorupnos/crow/api/internal/app"
	"github.com/doktorupnos/crow/api/internal/respond"
)

func TestHealthz(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(app.WithError(app.Healthz))
	defer server.Close()

	client := server.Client()
	req, err := http.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		server.URL,
		http.NoBody,
	)
	if err != nil {
		t.Fatal("NewRequest:", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Get() unexpected error:", err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			t.Log("error closing response body:", err)
		}
	}()

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

func TestError(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(app.WithError(app.Error))
	defer server.Close()

	client := server.Client()
	req, err := http.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		server.URL,
		http.NoBody,
	)
	if err != nil {
		t.Fatal("NewRequest:", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Get() unexpected error:", err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			t.Log("error closing response body:", err)
		}
	}()

	var v respond.ErrorResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&v)
	if err != nil {
		log.Fatal("json.Decode:", err)
	}

	t.Log(v.Message)
}
