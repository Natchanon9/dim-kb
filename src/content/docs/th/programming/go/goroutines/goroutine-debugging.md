---
title: "Debug Goroutine"
description: "race detector, stack dump และ goroutine profile"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

bug ด้าน concurrency ทำให้เกิดซ้ำได้ยาก ให้พึ่งเครื่องมือที่สังเกตโปรแกรมตอนรันจริง

## เครื่องมือ

| เครื่องมือ | หาอะไร |
| --- | --- |
| `go test -race`, `go run -race` | data race |
| `kill -QUIT <pid>` / Ctrl+\ | dump stack ของทุก goroutine แล้วจบโปรแกรม |
| `runtime.Stack(buf, true)` | dump stack จากในโปรแกรม |
| pprof `/debug/pprof/goroutine?debug=2` | ดู stack ของ goroutine ตอนรันผ่าน HTTP |
| `go tool trace` | timeline ของการจัดคิว การบล็อก และ GC |
| Delve (`dlv`) | debugger ไล่ทีละ goroutine |

```go title="stack.go"
buf := make([]byte, 1<<16)
n := runtime.Stack(buf, true) // true: ทุก goroutine
fmt.Println(string(buf[:n]))  // "goroutine 1 [running]: ..."
```

deadlock ที่ทุก goroutine ถูกบล็อกหมดจะจบด้วย `fatal error: all goroutines are asleep - deadlock!` พร้อม dump

:::caution[ข้อควรระวัง]
- race detector เจอ race เฉพาะในโค้ดที่ถูกรันจริง ต้องมี test ที่วิ่งผ่านส่วนที่ทำงานพร้อมกัน
:::
