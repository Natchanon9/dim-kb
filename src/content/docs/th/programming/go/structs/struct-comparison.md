---
title: "การเปรียบเทียบ Struct"
description: "== กับ struct, ใช้เป็น key ของ map และการเทียบ struct ที่มี slice"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

struct เทียบด้วย `==` ได้เมื่อทุก field เทียบได้ และ struct ที่เทียบได้ใช้เป็น key ของ map ได้ด้วย

## ตัวอย่าง

```go title="compare.go"
type Point struct{ X, Y int }

Point{1, 2} == Point{1, 2}           // true
seen := map[Point]bool{{1, 2}: true} // struct เป็น key ของ map

type Tagged struct {
	Name string
	Tags []string // field เป็น slice: เทียบไม่ได้
}
a := Tagged{"x", []string{"go"}}
b := Tagged{"x", []string{"go"}}
reflect.DeepEqual(a, b) // true ส่วน a == b จะ compile ไม่ผ่าน
```

## เทียบ struct ที่เทียบด้วย == ไม่ได้

- ใน test: `reflect.DeepEqual` หรือ `github.com/google/go-cmp`
- ใน production: เขียน method `Equal` เอง เร็วกว่าและบอกชัดว่าเท่ากันหมายถึงอะไร

:::caution[ข้อควรระวัง]
- เทียบ pointer (`&a == &b`) คือเทียบว่าเป็นตัวเดียวกัน ไม่ใช่เทียบเนื้อหา
- เพิ่ม field ที่เป็น slice หรือ map ทีหลังจะทำให้ทุก `==` ของ type นั้น compile ไม่ผ่าน
:::
