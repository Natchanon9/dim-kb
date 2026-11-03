---
title: "Enhanced ServeMux"
description: "Method patterns, wildcards and precedence in Go 1.22"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

Go 1.22 upgraded `http.ServeMux` with methods, `{wildcards}` and clear precedence, removing the main reason to add a router library.

## Example

```go title="routes.go"
mux.HandleFunc("GET /{$}", home)              // exactly "/"
mux.HandleFunc("GET /posts/{id}", showPost)   // r.PathValue("id")
mux.HandleFunc("GET /posts/latest", latest)   // more specific: wins over {id}

// /             → 200 home
// /posts/7      → 200 post 7
// /posts/latest → 200 latest
// /nope         → 404
```

## What changed from 1.21

| Before | Since 1.22 |
| --- | --- |
| Path only | `METHOD /path` |
| No parameters | `{name}`, `{rest...}` |
| Longest prefix wins | Most specific pattern wins |
| `/` matched everything | `/{$}` for exact root |

Use `r.SetPathValue` in tests to fake wildcards when calling a handler directly. Details on [ServeMux Routing](/dim-kb/programming/go/networking-http/servemux-routing/).

:::caution[Pitfalls]
- Old behaviour returns if `go.mod` says `go 1.21` or lower, or with `GODEBUG=httpmuxgo121=1`.
:::
