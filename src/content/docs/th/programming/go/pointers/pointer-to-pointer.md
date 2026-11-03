---
title: "Pointer to Pointer"
description: "**T ให้ function เปลี่ยน pointer ของผู้เรียก"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`**T` คือ pointer ไปยัง pointer ต้องใช้ก็ต่อเมื่อ function ต้องเปลี่ยนว่า pointer ของผู้เรียกชี้ไปที่ไหน

## ตัวอย่าง

```go title="pp.go"
func setNil(pp **int) { *pp = nil }

x := 1
p := &x
pp := &p
**pp = 2
fmt.Println(x, *pp == p) // 2 true
setNil(&p)
fmt.Println(p == nil) // true
```

ในโค้ด Go ทั่วไปแทบไม่ได้ใช้ การคืน pointer ตัวใหม่ (`p = rebuild(p)`) อ่านง่ายกว่า

:::caution[ข้อควรระวัง]
- ถ้ากำลังจะใช้ `**T` ลองดูก่อนว่าคืนค่ากลับไปจะง่ายกว่าไหม
:::
