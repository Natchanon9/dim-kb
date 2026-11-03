---
title: "What is Go"
description: "What makes Go stand out and what it is used for"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

Go is a compiled, statically typed, garbage-collected language from Google, open source since 2009. It was built so large server codebases compile fast and stay easy to read.

## Three core traits

| Trait | Meaning |
| --- | --- |
| **Compiled** | Builds to machine code; no VM at runtime |
| **Statically typed** | Type errors are caught at compile time |
| **Garbage-collected** | Memory is freed for you |

## What makes Go different

### Simple by design

25 keywords, `for` is the only loop, and there are no classes or inheritance. Structs and interfaces cover the same ground with less ceremony.

### Concurrency built in

A goroutine is a function running concurrently with the rest of the program. It starts with a stack of a few KB, so running thousands is normal. Channels pass data between goroutines.

### Batteries-included stdlib

HTTP, JSON, crypto and testing ship with the language. A JSON API needs no framework:

```go title="hello.go"
mux := http.NewServeMux()
mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "hello " + r.PathValue("name"),
	})
})
http.ListenAndServe(":8080", mux)
```

`GET /hello/gopher` returns `{"message":"hello gopher"}`. A `POST` gets `405` for free because the method is part of the pattern (Go 1.22+).

### One binary

`go build` produces a single self-contained executable. Cross-compiling is one env var: `GOOS=linux go build`.

## Explicit over clever

One code format (`gofmt`). Errors are plain values checked with `if err != nil`. No exceptions, operator overloading or macros. Code looks the same across teams, at the cost of a little more typing.

## Where Go fits

| Good fit | Poor fit |
| --- | --- |
| REST APIs, microservices | GUI and native mobile apps |
| CLIs and infra tooling (Docker, Kubernetes, Terraform) | ML and data science |
| Proxies and network services | Kernels, drivers, hard real-time |

:::caution[Pitfalls]
- Don't write Go like Java: no class hierarchies, no getter/setter for every field.
- Goroutines are cheap, not free. One that never exits is a memory leak.
- The binary is larger than C/C++ because it bundles the runtime and GC.
:::
