---
title: "JSON APIs"
description: "Decode requests, validate, and write JSON responses"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

A JSON endpoint does three things: decode the body, validate it, and write a JSON response with the right status.

## Example

```go title="users.go"
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
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
// {"name":"ann"} → 201 {"id":1,"name":"ann"}
// {}             → 400 {"error":"name required"}
```

## Habits

- Separate request/response structs from database models.
- Limit body size with `http.MaxBytesReader`.
- One consistent error shape for every endpoint.

:::caution[Pitfalls]
- Never return raw internal errors (SQL messages) to clients; log them and send a generic message.
:::
