---
title: "Profiling ด้วย pprof"
description: "หาว่าเวลา CPU และ memory ไปอยู่ที่ไหน"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

pprof สุ่มเก็บข้อมูลจากโปรแกรมที่กำลังรัน แล้วแสดงว่า function ไหนใช้ CPU จอง memory หรือถูกบล็อก profile ก่อนค่อย optimize

## ใน server

```go title="main.go"
import _ "net/http/pprof" // ลงทะเบียน /debug/pprof/ บน http.DefaultServeMux

go http.ListenAndServe("localhost:6060", nil) // อย่าเปิดบน port สาธารณะ
```

```sh
go tool pprof -http=:8081 http://localhost:6060/debug/pprof/profile?seconds=30  # CPU
go tool pprof http://localhost:6060/debug/pprof/heap                          # memory
```

## ในโค้ดหรือ test

```go title="cpu.go"
f, _ := os.Create("cpu.out")
pprof.StartCPUProfile(f)
busyWork()
pprof.StopCPUProfile()
```

```sh
go test -bench=. -cpuprofile=cpu.out -memprofile=mem.out
go tool pprof -http=:8081 cpu.out
```

| Profile | แสดง |
| --- | --- |
| profile (CPU) | เวลาไปหมดที่ไหน |
| heap / allocs | memory ที่ยังใช้อยู่ / จุดที่จอง |
| goroutine | stack ของทุก goroutine |
| block / mutex | การรอที่ channel / lock |

:::caution[ข้อควรระวัง]
- ห้ามเปิด `/debug/pprof` สู่สาธารณะ ทำให้ข้อมูลภายในรั่วและถูกใช้ทำ DoS ได้
:::
