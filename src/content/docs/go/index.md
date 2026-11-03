---
title: Go
description: Go section overview — reading order and sources
status: draft
sources:
  - https://algomaster.io/learn/go
  - https://mikelopster.dev/docs/c/goapi-essential/intro
---

Go notes limited to what backend work actually uses. Read in sidebar order.

| Part | Sections | Main source |
| --- | --- | --- |
| Language | Getting Started → Packages & Modules | [algomaster](https://algomaster.io/learn/go) |
| Concurrency | Concurrency, Generics | algomaster + mikelopster chapter 6 |
| Backend | HTTP & APIs → Deploy | [mikelopster GoAPI Essential](https://mikelopster.dev/docs/c/goapi-essential/intro) |

## Rules for pages in this section

- Summarise in my own words, never copy from the source, and always link the source at the end of the page.
- The real snippet lives in `examples/go/<section>/<page>/` and must pass `go test ./...`.
- Technical terms stay in English in the Thai translation (slice, goroutine, receiver) — see the [Glossary](/dim-kb/go/reference/glossary/).
