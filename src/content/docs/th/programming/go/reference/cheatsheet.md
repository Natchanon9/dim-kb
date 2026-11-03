---
title: "Go cheat sheet"
description: "syntax และคำสั่งที่ใช้บ่อยในหน้าเดียว"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

หน้าเดียวไว้ทวน syntax และคำสั่งที่ใช้บ่อยที่สุด แต่ละแถวมีหน้าเต็มอยู่ใน sidebar

## การประกาศ

| ต้องการ | เขียน |
| --- | --- |
| ตัวแปร | `var n int`, `n := 1` |
| constant / enum | `const Max = 10`, `const ( A = iota; B )` |
| function | `func add(a, b int) (int, error)` |
| method | `func (u *User) Rename(n string)` |
| struct | `` type User struct { Name string `json:"name"` } `` |
| interface | `type Store interface { Get(id int) (User, error) }` |
| generic func | `func Map[T, U any](s []T, f func(T) U) []U` |

## Collection

| ต้องการ | เขียน |
| --- | --- |
| slice | `s := []int{1, 2}`, `make([]int, 0, n)`, `s = append(s, 3)` |
| map | `m := map[string]int{}`, `v, ok := m[k]`, `delete(m, k)` |
| loop | `for i := range n`, `for i, v := range s`, `for k, v := range m` |
| sort | `slices.Sort(s)`, `slices.SortFunc(s, func(a, b T) int { ... })` |
| key ที่เรียงแล้ว | `slices.Sorted(maps.Keys(m))` (1.23+) |

## Error

| ต้องการ | เขียน |
| --- | --- |
| สร้าง | `errors.New("x")`, `fmt.Errorf("load %s: %w", p, err)` |
| เช็คค่า | `errors.Is(err, ErrNotFound)` |
| เช็ค type | `var ve *ValidationError; errors.As(err, &ve)` |
| หลายตัว | `errors.Join(e1, e2)` |

## Concurrency

| ต้องการ | เขียน |
| --- | --- |
| เริ่ม | `go f()` |
| รอ | `var wg sync.WaitGroup; wg.Add(1); defer wg.Done(); wg.Wait()` |
| channel | `ch := make(chan T, n)`, `ch <- v`, `v, ok := <-ch`, `close(ch)` |
| select + timeout | `select { case v := <-ch: ... case <-ctx.Done(): ... }` |
| lock | `mu.Lock(); defer mu.Unlock()` |
| ยกเลิก | `ctx, cancel := context.WithTimeout(ctx, 2*time.Second); defer cancel()` |

## HTTP

| ต้องการ | เขียน |
| --- | --- |
| route | `mux.HandleFunc("GET /users/{id}", h)`, `r.PathValue("id")` |
| JSON เข้า | `json.NewDecoder(r.Body).Decode(&req)` |
| JSON ออก | `w.Header().Set("Content-Type", "application/json"); json.NewEncoder(w).Encode(v)` |
| client | `&http.Client{Timeout: 5 * time.Second}` |

## คำสั่ง

```sh
go mod init example.com/app   go mod tidy
go run .                      go build -o app ./cmd/api
go test -race -cover ./...    go test -run TestX -v
go vet ./...                  gofmt -l .
go test -bench=. -benchmem    go tool pprof cpu.out
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w"
```

:::caution[ข้อควรระวัง]
- เช็คลิสต์ bug ที่เจอบ่อยฉบับเต็มอยู่ที่ [ข้อผิดพลาดที่พบบ่อย](/dim-kb/th/programming/go/best-practices/common-mistakes/)
:::
