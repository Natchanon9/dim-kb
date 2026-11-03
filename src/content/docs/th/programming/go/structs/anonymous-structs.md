---
title: "Anonymous Structs"
description: "struct type ใช้ครั้งเดียวโดยไม่ต้องตั้งชื่อ"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

anonymous struct ประกาศและใช้ในที่เดียว ไม่ต้องมี `type`

## ที่ใช้บ่อย

```go title="anon.go"
point := struct{ X, Y int }{1, 2}

tests := []struct { // test case แบบ table-driven
	in   string
	want int
}{
	{"a", 1},
	{"bb", 2},
}
fmt.Println(point, len(tests), tests[1].want) // {1 2} 2 2
```

ใช้สะดวกกับรูปแบบ JSON ของ request/response ที่ใช้ครั้งเดียวใน handler และใช้จัดกลุ่ม config ในตัวแปร

:::caution[ข้อควรระวัง]
- ถ้ารูปแบบเดียวกันโผล่สองที่ ให้ตั้งชื่อ type
:::
