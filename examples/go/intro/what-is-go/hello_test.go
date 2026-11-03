package whatisgo

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	srv := httptest.NewServer(NewMux())
	defer srv.Close()

	res, err := http.Get(srv.URL + "/hello/gopher")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := strings.TrimSpace(string(body)), `{"message":"hello gopher"}`; got != want {
		t.Errorf("body = %s, want %s", got, want)
	}
}

func TestWrongMethod(t *testing.T) {
	rec := httptest.NewRecorder()
	NewMux().ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/hello/gopher", nil))
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("POST status = %d, want 405", rec.Code)
	}
}
