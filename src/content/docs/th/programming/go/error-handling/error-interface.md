---
title: "error Interface"
description: "error เป็นค่าธรรมดาที่ implement Error() string"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`error` เป็น interface ในตัวภาษาที่มี method เดียว type ไหนมี `Error() string` ก็เป็น error ได้ และ `nil` หมายถึง "ไม่มี error"

## Interface

```go title="error.go"
// มีในตัวภาษา:
// type error interface { Error() string }

var err error = errors.New("disk full")
fmt.Println(err, err.Error() == "disk full") // disk full true
```

## Error เป็นค่า

Go ไม่มี exception function ที่อาจล้มเหลวจะคืน `error` เป็นผลลัพธ์ตัวสุดท้าย และผู้เรียกเช็คทันที:

```go
f, err := os.Open("config.yaml")
if err != nil {
	return err
}
```

เพราะ error เป็นค่า จึงเก็บ เปรียบเทียบ wrap และคืนค่าได้เหมือนค่าอื่นๆ

:::caution[ข้อควรระวัง]
- อย่าเทียบ **ข้อความ** ของ error (`err.Error() == "..."`) ข้อความเปลี่ยนได้ ให้ใช้ `errors.Is`/`errors.As`
:::
