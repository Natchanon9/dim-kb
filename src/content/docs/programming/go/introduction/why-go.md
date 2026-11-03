---
title: "Why Go?"
description: "What Go is good at, what it is bad at, and the trade-offs"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

Go trades expressiveness for simplicity, fast builds and easy deployment. That trade fits servers and tooling very well.

## Strengths

- **Fast feedback:** large projects compile in seconds.
- **Small language:** 25 keywords and one enforced format (`gofmt`), so codebases look alike.
- **Concurrency:** goroutines and channels are part of the language.
- **Deployment:** one static binary, cross-compiled with `GOOS`/`GOARCH`.
- **Stdlib:** production-ready HTTP server, JSON, crypto and testing.

## Trade-offs

| You give up | You get |
| --- | --- |
| Classes and inheritance | Structs, interfaces, composition |
| Exceptions | Errors as values, explicit `if err != nil` |
| Manual memory control | Garbage collection |
| Macros, operator overloading | Code anyone on the team can read |

## Where Go fits

| Good fit | Poor fit |
| --- | --- |
| REST/gRPC APIs, microservices | GUI and native mobile apps |
| CLIs and infra tooling (Docker, Kubernetes, Terraform) | ML and data science |
| Proxies, network services | Kernels, drivers, hard real-time |

:::caution[Pitfalls]
- Verbose error handling is by design. Wrapping it in helper magic usually makes code harder to read.
:::
