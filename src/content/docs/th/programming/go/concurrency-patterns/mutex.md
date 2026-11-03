---
title: "Mutex"
description: "ป้องกัน state ที่แชร์ด้วย sync.Mutex และ RWMutex"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

mutex ให้ goroutine เข้า critical section ได้ทีละตัว วาง mutex ไว้ข้างข้อมูลที่มันป้องกัน

## ตัวอย่าง

```go title="counter.go"
type SafeCounter struct {
	mu sync.Mutex
	m  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key]++
}
// goroutine 100 ตัวเรียก Inc("hits") → 100
```

## RWMutex

`sync.RWMutex` ให้อ่านได้หลายตัวพร้อมกัน หรือเขียนได้ตัวเดียว ใช้ `RLock`/`RUnlock` ตอนอ่าน ช่วยได้ก็ต่อเมื่ออ่านมากกว่าเขียนเยอะมาก ไม่อย่างนั้น `Mutex` ธรรมดาเร็วกว่า

:::caution[ข้อควรระวัง]
- อย่า copy struct ที่มี mutex ให้ใช้ pointer receiver
- `Lock` เข้าซ้ำไม่ได้: lock สองครั้งใน goroutine เดียวกันจะ deadlock
- ทำ critical section ให้สั้น อย่าทำ I/O ระหว่างถือ lock
:::
