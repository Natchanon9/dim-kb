---
title: "What is Go"
description: "What makes Go stand out and what it is used for"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
sources:
  - https://algomaster.io/learn/go/what-is-go
  - https://algomaster.io/learn/go/why-go
  - https://mikelopster.dev/docs/c/goapi-essential/chapter-1/intro
---

## TL;DR

- Go is a **compiled, statically typed, garbage-collected** language from Google (started by Griesemer, Pike and Thompson in 2007, open source since 2009)
- Designed to be **simple**: only 25 keywords, `for` is the only loop, no classes or inheritance — structs + interfaces instead
- **Concurrency is built into the language**: goroutines are very cheap (stacks start at a few KB) and talk to each other over channels
- Builds a **single binary** with no runtime or VM to install — easy to deploy, cross-compiles out of the box
- Strong stdlib: HTTP server, JSON, crypto and testing are production-ready

## Concept

Go came out of Google's frustration with C++ builds that took tens of minutes and Java-style code buried under layers of inheritance. The goal was a language that compiles fast, reads easily, and makes it painless to write servers that handle many requests at once.

The core philosophy is **explicit over clever**: there is one code format (`gofmt`), errors are plain values you check with `if err != nil` (no exceptions), and there is no operator overloading or macros. The result is that everyone's code looks alike, so you can open someone else's repo and follow it quickly — at the cost of typing a bit more.

Trade-offs worth knowing: generics only arrived in 1.18, the ecosystem is smaller than Java/Python, the GC gives less memory control than Rust/C++, and Go is a poor fit for GUIs, native mobile, and ML/data science.

## Code

The smallest useful API with only the standard library, no framework (Go 1.22+ lets you put the method and path params in the pattern).

```go title="hello.go"
mux := http.NewServeMux()
mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "hello " + r.PathValue("name"),
	})
})
http.ListenAndServe(":8080", mux)
// GET  /hello/gopher -> {"message":"hello gopher"}
// POST /hello/gopher -> 405 Method Not Allowed (free from the pattern)
```

Full version + tests: `examples/go/intro/what-is-go/`

## Gotchas

:::caution
- "Simple" ≠ "can do less" — it means there are only a few ways to do things. Don't write Go like Java (class hierarchies, a getter/setter for every field).
- Goroutines are cheap, not free: leaving them stuck = a goroutine leak that keeps eating memory (see the Concurrency section).
- The single binary is bigger than C/C++ because it bundles the runtime and GC.
:::

## When to use

Reach for Go for REST APIs, microservices, CLIs and infra tooling (Docker, Kubernetes and Terraform are all written in Go) — especially when you need to handle many concurrent requests and want to ship a small container.

## References

- <https://algomaster.io/learn/go/what-is-go>
- <https://algomaster.io/learn/go/why-go>
- <https://mikelopster.dev/docs/c/goapi-essential/chapter-1/intro>
