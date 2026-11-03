---
title: "Range"
description: "วนลูป slice, string, map, channel และ integer ด้วย range"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`for ... range` ไล่ไปตามข้อมูลทีละตัว แต่ละรอบได้ index (หรือ key) และค่า

## range คืนค่าอะไร

| วนบน | ค่าแรก | ค่าที่สอง |
| --- | --- | --- |
| Slice / array | index | สำเนาของ element |
| String | ตำแหน่ง byte | rune (ถอด UTF-8 แล้ว) |
| Map | key | value (ลำดับสุ่ม) |
| Channel | element | (ไม่มี) |
| Integer `n` (1.22+) | 0 … n-1 | (ไม่มี) |

```go title="range.go"
for i, v := range []string{"a", "b"} {
	fmt.Println(i, v) // 0 a, 1 b
}
for i, r := range "gö" {
	fmt.Println(i, string(r)) // 0 g, 1 ö
}
```

ตัดส่วนที่ไม่ใช้ทิ้งได้: `for _, v := range s`, `for i := range s`, `for range ticker.C`

:::caution[ข้อควรระวัง]
- `v` เป็นสำเนา `v.Count++` บน slice ของ struct ไม่เปลี่ยนอะไร ต้องใช้ `s[i].Count++`
- ลำดับการวน map สุ่มโดยตั้งใจ ถ้าลำดับสำคัญให้ sort key ก่อน
:::
