---
title: "First-Class Functions"
description: "function เป็นค่า: เก็บในตัวแปร ส่งเป็น parameter เก็บใน map และตั้งเป็น type"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

function เป็นค่าธรรมดา เก็บในตัวแปร ส่งเป็น argument คืนค่า และเก็บใน map ได้

## ส่ง function เป็น argument

```go title="apply.go"
func apply(nums []int, f func(int) int) []int {
	out := make([]int, 0, len(nums))
	for _, n := range nums {
		out = append(out, f(n))
	}
	return out
}

square := func(n int) int { return n * n }
apply([]int{1, 2, 3}, square) // [1 4 9]
```

## ตารางเลือก function

```go title="ops.go"
ops := map[string]func(a, b int) int{
	"add": func(a, b int) int { return a + b },
	"mul": func(a, b int) int { return a * b },
}
ops["mul"](3, 4) // 12
```

signature ที่ใช้บ่อยตั้งชื่อเป็น type ได้: `type HandlerFunc func(http.ResponseWriter, *http.Request)`

:::caution[ข้อควรระวัง]
- zero value ของ func คือ `nil` เรียกแล้ว panic
- func เทียบได้กับ `nil` เท่านั้น เลยเป็น key ของ map ไม่ได้
:::
