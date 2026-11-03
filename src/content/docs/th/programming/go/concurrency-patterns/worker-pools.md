---
title: "Worker Pools"
description: "goroutine จำนวนคงที่ประมวลผลคิวของงาน"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

worker pool เริ่ม goroutine N ตัวที่ดึงงานจาก channel เดียวกัน เป็นการจำกัดจำนวนงานพร้อมกัน ช่วยปกป้อง database และ API ไม่ให้รับภาระเกิน

## ตัวอย่าง

```go title="pool.go"
func squareAll(jobs []int, workers int) []int {
	in := make(chan int)
	out := make(chan int, len(jobs))
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range in {
				out <- j * j
			}
		}()
	}
	for _, j := range jobs {
		in <- j
	}
	close(in) // worker ออกจาก range loop
	wg.Wait()
	close(out)
	// เก็บจาก out → [1 4 9 16]
}
```

## จำนวน worker

- งานที่ใช้ CPU: ประมาณ `runtime.NumCPU()` ตัว
- งานที่รอ I/O: มากกว่านั้น จำกัดตามที่ service ปลายทางรับไหว

:::caution[ข้อควรระวัง]
- ปิด channel ของงานครั้งเดียวหลังส่งตัวสุดท้าย และปิด channel ผลลัพธ์หลังทุก worker เสร็จเท่านั้น
:::
