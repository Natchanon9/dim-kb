---
title: "พื้นฐาน Goroutine"
description: "เริ่มงานแบบ concurrent ด้วย keyword go"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

goroutine คือ function ที่รันไปพร้อมกับส่วนอื่นของโปรแกรม ใส่ `go` หน้าการเรียก function เพื่อเริ่ม แล้วมันจะ return ทันที

## ตัวอย่าง

```go title="basics.go"
var wg sync.WaitGroup
results := make([]int, 3)
for i := range 3 {
	wg.Add(1)
	go func() {
		defer wg.Done()
		results[i] = i * i // แต่ละ goroutine เขียน index ของตัวเอง
	}()
}
wg.Wait()
fmt.Println(results) // [0 1 4]
```

## เรื่องที่ต้องรู้

- statement `go` ไม่คืนค่า ส่งผลลัพธ์ผ่าน channel หรือเขียนลง memory ที่เป็นของมันเอง
- ต้อง **รอ** goroutine เอง (WaitGroup, channel, errgroup) ไม่มีอะไรรอให้อัตโนมัติ
- goroutine ถูกมาก: เริ่มด้วย stack เล็ก (ตอนนี้ 2 KB) แล้วขยายตามต้องการ

:::caution[ข้อควรระวัง]
- สอง goroutine เขียนตัวแปรเดียวกันโดยไม่ lock คือ data race ให้รัน test ด้วย `-race`
- `go f(x)` ประเมิน `x` ทันทีใน goroutine ที่เรียก
:::
