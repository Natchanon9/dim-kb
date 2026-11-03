---
title: "Error Handling Best Practices"
description: "เช็คลิสต์ error สำหรับโค้ด production"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

error handling ที่ดีทำให้ไล่ปัญหาง่าย และตอบสนองต่อความผิดพลาดได้ง่าย

## เช็คลิสต์

- **wrap พร้อม context** ทุกชั้น: `fmt.Errorf("charge order %d: %w", id, err)`
- **จัดการครั้งเดียว:** log **หรือ** return ไม่ทำทั้งคู่
- **เทียบด้วย `errors.Is`/`errors.As`** ไม่เทียบข้อความ
- **sentinel สำหรับ "ตัวไหน" type สำหรับ "รายละเอียด"**
- **ข้อความตัวพิมพ์เล็ก** ไม่มีเครื่องหมายวรรคตอนท้าย
- **อย่าให้รายละเอียดภายในหลุด** ไปถึง client ของ API: แปลง error เป็นข้อความที่ปลอดภัยและ status code ที่ชั้นนอกสุด
- **รายงานหลายปัญหาพร้อมกัน** ด้วย `errors.Join` (Go 1.20+):

```go title="join.go"
err := errors.Join(
	errors.New("name required"),
	errors.New("email invalid"),
)
fmt.Println(err)
// name required
// email invalid
```

`errors.Is` และ `errors.As` ค้นในทุก error ที่ join ไว้

:::caution[ข้อควรระวัง]
- panic กับความล้มเหลวที่คาดได้ และทิ้ง error ด้วย `_` คือสองเรื่องที่โดนทักใน review โค้ด Go บ่อยที่สุด
:::
