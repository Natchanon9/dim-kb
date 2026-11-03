---
title: "Fan-Out"
description: "หลาย worker อ่านจาก channel เดียวเพื่อแบ่งงาน"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

fan-out เริ่ม goroutine หลายตัวที่อ่านจาก channel เดียวกัน แต่ละค่าไปถึง worker แค่ตัวเดียว งานจึงกระจายไปทั่ว

## ตัวอย่าง

```go title="fanout.go"
jobs := emit(1, 2, 3, 4)
results := make(chan int)
var wg sync.WaitGroup
for range 3 { // worker 3 ตัว
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := range jobs {
			results <- j * 10
		}
	}()
}
go func() { wg.Wait(); close(results) }()
// เก็บผลลัพธ์ → [10 20 30 40] หลัง sort
```

fan-out กับ fan-in รวมกันคือรูปแบบคลาสสิกของการประมวลผลขนาน ถ้าจำนวน worker คงที่ก็คือ worker pool

:::caution[ข้อควรระวัง]
- ผลลัพธ์มาลำดับไหนก็ได้ ถ้าต้องจับคู่กลับให้แนบ index หรือ ID ไปด้วย
:::
