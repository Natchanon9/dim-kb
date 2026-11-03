---
title: "httptest"
description: "test HTTP handler และ client โดยไม่ต้องใช้ network จริง"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`net/http/httptest` มี response writer ปลอมไว้ test handler และมี server จริงในเครื่องไว้ test client

## Test handler

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

## Test client

```go
srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"id":1}`))
}))
defer srv.Close()
client := NewAPIClient(srv.URL) // ชี้ client ไปที่ server ปลอม
```

ส่ง router ทั้งตัว (`mux.ServeHTTP(rec, req)`) เพื่อ test routing และ middleware ไปด้วย

:::caution[ข้อควรระวัง]
- `NewRequest` panic ถ้า URL ผิด เพราะออกแบบมาใช้ใน test เท่านั้น
:::
