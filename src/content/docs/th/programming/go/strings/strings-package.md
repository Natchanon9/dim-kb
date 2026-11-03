---
title: "strings Package"
description: "ค้นหา แยก ต่อ ตัดช่องว่าง และแทนที่ข้อความ"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`strings` รวม helper สำหรับข้อความที่ใช้ทุกวัน ทุก function คืน string ใหม่

## Function ที่ใช้บ่อย

```go title="strings.go"
t := strings.TrimSpace("  Hello, Go  ")
strings.ToUpper(t)                      // HELLO, GO
strings.Contains(t, "Go")               // true
strings.HasPrefix(t, "He")              // true
strings.Index(t, "Go")                  // 7
strings.Split("a,b,c", ",")             // [a b c]
strings.Join([]string{"x", "y"}, "-")   // x-y
strings.ReplaceAll("aaa", "a", "b")     // bbb
strings.Fields(" a  b ")                // [a b]
strings.EqualFold("Go", "GO")           // true (ไม่สนตัวพิมพ์)
key, value, found := strings.Cut("key=value", "=") // key value true
```

`Cut` (Go 1.18+) ใช้แทนโค้ด `Index` + ตัด string แบบ "แยกครั้งเดียว" ได้เกือบหมด

:::caution[ข้อควรระวัง]
- `Split("", ",")` คืน `[""]` (มี element ว่าง 1 ตัว) ไม่ใช่ slice ว่าง
- เทียบด้วย `ToLower(a) == ToLower(b)` จอง memory สองรอบ ให้ใช้ `EqualFold`
:::
