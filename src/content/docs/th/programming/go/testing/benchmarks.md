---
title: "Benchmarks"
description: "วัดความเร็วและการจอง memory ด้วย BenchmarkXxx"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

benchmark คือ `BenchmarkXxx(b *testing.B)` framework จะเลือกจำนวนรอบเองจนได้เวลาต่อ operation ที่นิ่ง

## ตัวอย่าง

```go title="join_test.go"
func BenchmarkJoin(b *testing.B) {
	parts := []string{"a", "b", "c"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ { // Go 1.24+: for b.Loop() { ... }
		strings.Join(parts, ",")
	}
}
```

```sh
go test -bench=. -benchmem ./...
# BenchmarkJoin-16   12000000   95 ns/op   8 B/op   1 allocs/op
```

## เปรียบเทียบการเปลี่ยนแปลง

รันหลายรอบ (`-count=10`) ทั้งก่อนและหลังแก้ แล้วเทียบด้วย `benchstat old.txt new.txt` การรันครั้งเดียวคลาดเคลื่อนมาก

:::caution[ข้อควรระวัง]
- setup ที่อยู่ใน loop จะถูกวัดด้วย ให้ทำก่อน loop หรือเรียก `b.ResetTimer()`
- compiler อาจตัดงานที่ไม่ได้ใช้ผลลัพธ์ทิ้ง `b.Loop()` (1.24+) ป้องกันเรื่องนี้
:::
