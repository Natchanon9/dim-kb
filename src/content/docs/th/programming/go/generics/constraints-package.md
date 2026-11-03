---
title: "constraints Package"
description: "comparable, cmp.Ordered และ x/exp/constraints แบบเก่า"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

constraint สำเร็จรูปที่ต้องใช้มีน้อยมาก ตั้งแต่ Go 1.21 `cmp.Ordered` ใน standard library แทนการใช้ `golang.org/x/exp/constraints` ได้เกือบหมด

## ตัวอย่าง

```go title="keys.go"
func Index[T comparable](s []T, v T) int { // comparable: == และ !=
	for i, x := range s {
		if x == v {
			return i
		}
	}
	return -1
}

func SortedKeys[K cmp.Ordered, V any](m map[K]V) []K { // cmp.Ordered: < >
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
// Index([]string{"a","b"}, "b") → 1 ; SortedKeys → [a b]
```

| ต้องการ | ใช้ |
| --- | --- |
| เทียบเท่ากัน, key ของ map | `comparable` (มีในตัว) |
| เรียงลำดับ | `cmp.Ordered` |
| integer อย่างเดียว, float อย่างเดียว | สร้างเองด้วย `~int \| ~int64 ...` |

:::caution[ข้อควรระวัง]
- `x/exp/constraints` ยังเป็นแบบทดลอง โค้ดใหม่ให้ใช้ `cmp`
:::
