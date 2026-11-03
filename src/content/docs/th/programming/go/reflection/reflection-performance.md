---
title: "Performance ของ Reflection"
description: "reflection แพงแค่ไหน และลดต้นทุนอย่างไร"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

reflection ช้ากว่าโค้ดตรงๆ มาก: จอง memory, ค้นหาตามชื่อ และทำให้ compiler optimize ไม่ได้

## Benchmark

```go title="bench_test.go"
func BenchmarkDirectField(b *testing.B) {
	p := point{1, 2}
	for i := 0; i < b.N; i++ {
		sink = p.X
	}
}

func BenchmarkReflectField(b *testing.B) {
	p := point{1, 2}
	for i := 0; i < b.N; i++ {
		sink = int(reflect.ValueOf(p).FieldByName("X").Int())
	}
}
// DirectField   ~0.2 ns/op
// ReflectField  ~35 ns/op   (ช้ากว่าประมาณ 100 เท่าบนเครื่องนี้)
```

## ทำให้ถูกลง

- **cache** งานต่อ type (index ของ field, tag ที่ parse แล้ว) ใน `sync.Map` ที่ใช้ `reflect.Type` เป็น key แบบที่ `encoding/json` ทำ
- ใช้ `Field(i)` กับ index ที่ cache ไว้แทน `FieldByName`
- generate โค้ดแทน (`go generate`, `sqlc`, `easyjson`) สำหรับส่วนที่ทำงานหนัก

:::caution[ข้อควรระวัง]
- วัดก่อนปรับ request ที่ทำ I/O อยู่แล้ว 35 ns แทบไม่มีผล
:::
