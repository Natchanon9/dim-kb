---
title: "ServeMux Routing"
description: "Method and path-parameter patterns in Go 1.22+"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

Since Go 1.22, `http.ServeMux` patterns can include the HTTP method and `{name}` wildcards, which covers most routing needs without a third-party router.

## Example

```go title="routes.go"
mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "user ", r.PathValue("id"))
})
mux.HandleFunc("POST /users", createUser)
mux.HandleFunc("GET /files/{path...}", serveFile) // {x...} matches the rest

// GET /users/42      → 200 user 42
// POST /users        → 201 created
// GET /files/a/b.txt → 200 file a/b.txt
// DELETE /users/42   → 405 Method Not Allowed
```

## Matching rules

- The **most specific** pattern wins, regardless of registration order.
- `GET` also matches `HEAD`.
- A trailing `/` (`/static/`) matches the whole subtree; `{$}` (`/{$}`) matches the exact path only.
- Conflicting patterns panic at registration.

:::caution[Pitfalls]
- `go.mod` must say `go 1.22` or later; otherwise the old matching rules apply and `GET /x` is treated as a literal path.
:::
