---
title: "Fan-In"
description: "รวมหลาย channel เป็นหนึ่งเดียว"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

fan-in รวมค่าจาก channel ขาเข้าหลายตัวไว้ใน channel ขาออกตัวเดียว ใช้ goroutine หนึ่งตัวต่อหนึ่ง input เพื่อ copy ค่า และมี goroutine อีกตัวปิด output หลังทุก input เสร็จ

## ตัวอย่าง

```go title="merge.go"
func merge(chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, c := range chs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}()
	}
	go func() { wg.Wait(); close(out) }()
	return out
}

for v := range merge(emit(1, 2), emit(3)) { ... } // 1, 2, 3 ลำดับไม่แน่นอน
```

ใช้รวมผลลัพธ์จาก worker ที่ทำงานขนาน หรือรวม event จากหลายแหล่ง

:::caution[ข้อควรระวัง]
- ลำดับข้าม input ไม่ถูกรักษาไว้
- ถ้าฝั่งที่รับอาจหยุดก่อน ให้เพิ่ม case `ctx.Done()` ไม่อย่างนั้น goroutine ที่ copy ค่าจะ leak
:::
