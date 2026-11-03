---
title: "Go cheat sheet"
description: "Common syntax and commands on one page"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

One-page reminder of the syntax and commands used most. Each row links back to the full topic through the sidebar.

## Declarations

| Want | Write |
| --- | --- |
| Variable | `var n int`, `n := 1` |
| Constant / enum | `const Max = 10`, `const ( A = iota; B )` |
| Function | `func add(a, b int) (int, error)` |
| Method | `func (u *User) Rename(n string)` |
| Struct | `` type User struct { Name string `json:"name"` } `` |
| Interface | `type Store interface { Get(id int) (User, error) }` |
| Generic func | `func Map[T, U any](s []T, f func(T) U) []U` |

## Collections

| Want | Write |
| --- | --- |
| Slice | `s := []int{1, 2}`, `make([]int, 0, n)`, `s = append(s, 3)` |
| Map | `m := map[string]int{}`, `v, ok := m[k]`, `delete(m, k)` |
| Loop | `for i := range n`, `for i, v := range s`, `for k, v := range m` |
| Sort | `slices.Sort(s)`, `slices.SortFunc(s, func(a, b T) int { ... })` |
| Sorted keys | `slices.Sorted(maps.Keys(m))` (1.23+) |

## Errors

| Want | Write |
| --- | --- |
| Create | `errors.New("x")`, `fmt.Errorf("load %s: %w", p, err)` |
| Check value | `errors.Is(err, ErrNotFound)` |
| Check type | `var ve *ValidationError; errors.As(err, &ve)` |
| Several | `errors.Join(e1, e2)` |

## Concurrency

| Want | Write |
| --- | --- |
| Start | `go f()` |
| Wait | `var wg sync.WaitGroup; wg.Add(1); defer wg.Done(); wg.Wait()` |
| Channel | `ch := make(chan T, n)`, `ch <- v`, `v, ok := <-ch`, `close(ch)` |
| Select + timeout | `select { case v := <-ch: ... case <-ctx.Done(): ... }` |
| Lock | `mu.Lock(); defer mu.Unlock()` |
| Cancel | `ctx, cancel := context.WithTimeout(ctx, 2*time.Second); defer cancel()` |

## HTTP

| Want | Write |
| --- | --- |
| Route | `mux.HandleFunc("GET /users/{id}", h)`, `r.PathValue("id")` |
| JSON in | `json.NewDecoder(r.Body).Decode(&req)` |
| JSON out | `w.Header().Set("Content-Type", "application/json"); json.NewEncoder(w).Encode(v)` |
| Client | `&http.Client{Timeout: 5 * time.Second}` |

## Commands

```sh
go mod init example.com/app   go mod tidy
go run .                      go build -o app ./cmd/api
go test -race -cover ./...    go test -run TestX -v
go vet ./...                  gofmt -l .
go test -bench=. -benchmem    go tool pprof cpu.out
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w"
```

:::caution[Pitfalls]
- The full checklist of common bugs is on [Common Mistakes](/dim-kb/programming/go/best-practices/common-mistakes/).
:::
