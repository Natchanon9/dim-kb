---
title: "Sizeof, Alignof, Offsetof"
description: "รูปแบบของ struct, padding และลำดับ field"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

field ถูกจัดให้ตรงตามขนาดของมัน compiler จึงเติม padding ให้ ลำดับของ field เปลี่ยนขนาดของ struct ได้

## ตัวอย่าง

```go title="layout.go"
type padded struct {
	A bool  // 1 + padding 7
	B int64 // 8
	C bool  // 1 + padding 7
}

type packed struct {
	B int64 // 8
	A bool  // 1
	C bool  // 1 + padding 6
}

unsafe.Sizeof(padded{}), unsafe.Sizeof(packed{}) // 24 16
unsafe.Alignof(int64(0))                         // 8
unsafe.Offsetof(padded{}.B)                      // 8
```

เรียง field จากใหญ่ไปเล็กมักลด padding ได้มากที่สุด สำคัญกับ struct ที่เก็บเป็นล้านตัว (entry ของ cache, slice ของ record)

analyzer `fieldalignment` (golang.org/x/tools) บอกได้ว่า struct ไหนเล็กลงได้

:::caution[ข้อควรระวัง]
- อย่าสลับ field เพื่อลดขนาดจนอ่านยาก เว้นแต่ profile แล้วเห็นว่า memory ตึงจริง
:::
