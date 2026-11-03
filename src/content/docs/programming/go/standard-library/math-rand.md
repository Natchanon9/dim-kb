---
title: "math/rand"
description: "Pseudo-random numbers with math/rand/v2"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

`math/rand/v2` (Go 1.22+) generates fast pseudo-random numbers for simulations, jitter, sampling and shuffles. The top-level functions are seeded automatically.

## Example

```go title="rand.go"
rand.IntN(10)           // 0..9
rand.Float64()          // [0.0, 1.0)
rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
rand.N(5 * time.Second) // random duration, e.g. retry jitter

a := rand.New(rand.NewPCG(1, 2)) // fixed seed: same sequence every run
b := rand.New(rand.NewPCG(1, 2))
a.IntN(100) == b.IntN(100) // true
```

## v1 vs v2

| `math/rand` (v1) | `math/rand/v2` |
| --- | --- |
| `Intn` | `IntN` |
| Needed `rand.Seed` before Go 1.20 | Always auto-seeded |
| One source type | PCG, ChaCha8 |

:::caution[Pitfalls]
- Not for passwords, tokens or keys; use `crypto/rand`.
- A `*rand.Rand` with its own source isn't safe for concurrent use; the top-level functions are.
:::
