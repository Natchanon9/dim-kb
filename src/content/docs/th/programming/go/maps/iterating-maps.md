---
title: "วนลูป Map"
description: "range บน map และทำให้ลำดับคงที่"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`for k, v := range m` ไล่ครบทุก entry ในลำดับ **สุ่ม** ที่ตั้งใจให้เปลี่ยนทุกครั้งที่รัน

## วนแบบเรียงลำดับ

```go title="sorted.go"
m := map[string]int{"b": 2, "a": 1, "c": 3}
keys := make([]string, 0, len(m))
for k := range m {
	keys = append(keys, k)
}
slices.Sort(keys)
for _, k := range keys {
	fmt.Println(k, m[k]) // a 1, b 2, c 3
}
```

ตั้งแต่ Go 1.23 เขียนบรรทัดเดียวได้: `slices.Sorted(maps.Keys(m))`

## แก้ map ระหว่าง range

ลบ entry ระหว่าง `range` ได้อย่างปลอดภัย ส่วน entry ที่เพิ่มระหว่าง `range` อาจถูกวนถึงหรือไม่ก็ได้

:::caution[ข้อควรระวัง]
- อย่าพึ่งลำดับของ map ใน output, test หรือ JSON array ที่สร้างจาก map
:::
