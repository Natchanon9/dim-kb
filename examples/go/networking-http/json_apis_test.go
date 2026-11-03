package networkinghttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) // headers first, then status, then body
	json.NewEncoder(w).Encode(v)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
		return
	}
	if req.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "name required"})
		return
	}
	writeJSON(w, http.StatusCreated, map[string]any{"id": 1, "name": req.Name})
}

func Example_jsonAPI() {
	for _, body := range []string{`{"name":"ann"}`, `{}`} {
		rec := httptest.NewRecorder()
		createUser(rec, httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(body)))
		fmt.Print(rec.Code, " ", rec.Body.String())
	}
	// Output:
	// 201 {"id":1,"name":"ann"}
	// 400 {"error":"name required"}
}
