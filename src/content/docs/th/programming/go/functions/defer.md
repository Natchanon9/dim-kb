---
title: "defer"
description: "รัน cleanup ตอน function return เรียงแบบ LIFO"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`defer` นัดให้เรียก function ตอนที่ function ที่ครอบอยู่ return แม้จะ panic ก็ตาม ใช้วาง cleanup ไว้ติดกับตอน setup

## ลำดับและการประเมิน argument

```go title="defer.go"
for i := range 3 {
	defer fmt.Print(" ", i) // เข้าทีหลัง ออกก่อน
}
fmt.Print("start") // พิมพ์: start 2 1 0

x := 1
defer fmt.Println("deferred:", x) // ประเมิน x ทันที → 1
x = 2
```

## Cleanup ที่เจอบ่อย

```go
f, err := os.Open(path)
if err != nil {
	return err
}
defer f.Close()
```

ใช้แบบเดียวกันกับ `mu.Unlock()`, `resp.Body.Close()`, `rows.Close()`, `cancel()`

:::caution[ข้อควรระวัง]
- `defer` ใน loop รันตอน **function** return เท่านั้น ถ้า defer `Close` เป็นพันครั้ง ไฟล์จะเปิดค้าง ให้ย้าย body ของ loop ไปเป็น function
- argument ถูกประเมินทันที ถ้าต้องการค่าล่าสุดให้ห่อด้วย `func() { ... }()`
:::
