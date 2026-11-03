---
title: "ข้อผิดพลาดที่พบบ่อย"
description: "bug ที่โผล่ซ้ำแล้วซ้ำอีกในโค้ด Go"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

ส่วนใหญ่มีหน้าของตัวเองอยู่แล้ว หน้านี้เป็นเช็คลิสต์ไว้ไล่ตอน review

## สองเรื่องคลาสสิก

```go title="mistakes.go"
m := map[string]int{} // ไม่ใช่ `var m map[string]int`: เขียนลง nil map จะ panic
m["a"]++

func find(ok bool) error {
	if !ok {
		return &notFound{}
	}
	return nil // nil ตรงๆ ห้ามเป็น pointer ที่เป็น nil
}
```

## เช็คลิสต์

| ข้อผิดพลาด | วิธีแก้ |
| --- | --- |
| เขียนลง nil map | `make` ก่อน |
| คืน typed nil เป็น `error` | `return nil` |
| `defer` ใน loop ยาว | ย้าย body ไปเป็น function |
| ทิ้ง error ด้วย `_` | จัดการ หรือให้เหตุผลชัดเจน |
| goroutine ที่ไม่มีทางออก | context หรือ done channel |
| `http.Get` / client ไม่มี timeout | `&http.Client{Timeout: ...}` |
| ไม่ปิด `resp.Body` / `rows` | `defer x.Close()` |
| แก้สำเนาจาก `range` | index เข้าไปใน slice |
| แชร์ backing array หลัง `append` | copy เมื่อไม่ชัดว่าใครเป็นเจ้าของ |
| data race | `go test -race` |
| เทียบ error ด้วย `==` | `errors.Is` / `errors.As` |

:::caution[ข้อควรระวัง]
- เรื่องส่วนใหญ่ข้างบน `go vet`, staticcheck และ `-race` จับได้ ให้รันใน CI
:::
