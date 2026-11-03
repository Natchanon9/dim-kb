---
title: "Nil Pointers"
description: "เลี่ยง panic จาก nil dereference และเขียน method ที่รับ nil ได้"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

nil pointer ไม่ได้ชี้ไปไหน อ่าน field ผ่านมันจะ panic แต่เรียก method บนมันได้ method จึงจัดการกรณี nil เองได้

## Method ที่รับ nil ได้

```go title="node.go"
type Node struct {
	Val  int
	Next *Node
}

func (n *Node) Sum() int {
	if n == nil { // receiver เป็น nil ได้ ให้กันไว้
		return 0
	}
	return n.Val + n.Next.Sum()
}

list := &Node{1, &Node{2, nil}}
var empty *Node
fmt.Println(list.Sum(), empty.Sum()) // 3 0
```

## nil pointer มาจากไหน

- field ที่เป็น pointer ใน struct ที่ยังไม่ได้กำหนดค่า
- function คืน `nil, err` แล้วผู้เรียกไม่สน `err`
- ค้น `map[K]*V` ด้วย key ที่ไม่มี

:::caution[ข้อควรระวัง]
- เช็ค `err` ก่อนใช้ผลลัพธ์ตัวอื่นเสมอ nil dereference ส่วนใหญ่เริ่มจากตรงนี้
:::
