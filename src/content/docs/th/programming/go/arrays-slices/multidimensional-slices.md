---
title: "Slice หลายมิติ"
description: "slice ของ slice สำหรับตารางและ matrix"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

Go ไม่มี slice 2 มิติในตัว ตารางคือ slice ของแถว และทุกแถวเป็น slice แยกที่ต้องสร้างเอง

## สร้างตาราง

```go title="grid.go"
grid := make([][]int, 2)
for i := range grid {
	grid[i] = make([]int, 3) // แต่ละแถวเป็น slice ของตัวเอง
}
grid[1][2] = 9
fmt.Println(grid) // [[0 0 0] [0 0 9]]
```

แต่ละแถวยาวไม่เท่ากันได้ (jagged) เหมาะกับอย่าง adjacency list

## ขนาดคงที่

ถ้าทั้งสองมิติขนาดตายตัว ใช้ array ซ้อน array ง่ายกว่า: `var board [8][8]int`

:::caution[ข้อควรระวัง]
- `make([][]int, 2)` ได้แถวที่เป็น **nil** 2 แถว `grid[0][0]` จะ panic จนกว่าจะสร้างแต่ละแถว
:::
