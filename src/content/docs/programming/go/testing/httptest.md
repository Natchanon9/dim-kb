---
title: "httptest"
description: "Test HTTP handlers and clients without a real network"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`net/http/httptest` gives you a fake response writer for handler tests and a real local server for client tests.

## Testing a handler

```go title="health_test.go"
func TestHealth(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	healthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if got := rec.Body.String(); got != `{"status":"ok"}` {
		t.Errorf("body = %s", got)
	}
}
```

## Testing a client

```go
srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"id":1}`))
}))
defer srv.Close()
client := NewAPIClient(srv.URL) // point your client at the fake server
```

Pass the full router (`mux.ServeHTTP(rec, req)`) to test routing and middleware too.

:::caution[Pitfalls]
- `NewRequest` panics on a bad URL; it's meant for tests only.
:::
