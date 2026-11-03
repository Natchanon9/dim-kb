---
title: "Effective Go"
description: "สำนวนหลักที่ทำให้โค้ด Go ดูเป็น Go"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

"Effective Go" (go.dev/doc/effective_go) และ wiki Go Code Review Comments กำหนดสไตล์ของภาษา สำนวนไม่กี่ข้อครอบคลุมเกือบทั้งหมด

## ทำให้ zero value ใช้ได้

```go title="buffer.go"
type Buffer struct{ data []byte }

func (b *Buffer) Write(p []byte) { b.data = append(b.data, p...) }

var b Buffer // ไม่ต้องมี constructor
b.Write([]byte("go"))
```

## สำนวนหลัก

- **gofmt ทุกอย่าง** ไม่ต้องเถียงเรื่องสไตล์
- **จัดการ error ทันที** return ก่อน และให้ทางหลักไม่ย่อหน้าลึก
- **ชื่อสั้นสำหรับ scope สั้น** และ getter ไม่ต้องมี `Get`
- **รับ interface คืน struct** และทำ interface ให้เล็ก
- **แชร์ memory ด้วยการสื่อสาร** ตรงที่ทำให้โค้ดอ่านง่ายขึ้น
- **ชัดดีกว่าฉลาด** copy โค้ดนิดหน่อยดีกว่าเพิ่ม dependency นิดหน่อย

:::caution[ข้อควรระวัง]
- ย้าย pattern จาก Java/C# มาตรงๆ (inheritance ลึก, factory ทุกอย่าง, getter/setter) จะได้ Go ที่ไม่เป็นสำนวน Go
:::
