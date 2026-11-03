package introduction

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "hello " + r.PathValue("name"),
		})
	})
	return mux
}

func TestHelloAPI(t *testing.T) {
	srv := httptest.NewServer(newMux())
	defer srv.Close()

	res, err := http.Get(srv.URL + "/hello/gopher")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	if got := strings.TrimSpace(string(body)); got != `{"message":"hello gopher"}` {
		t.Errorf("body = %s", got)
	}
}

func TestHelloWrongMethod(t *testing.T) {
	rec := httptest.NewRecorder()
	newMux().ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/hello/gopher", nil))
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("POST status = %d, want 405", rec.Code)
	}
}
