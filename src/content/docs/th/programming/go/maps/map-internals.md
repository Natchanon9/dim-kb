---
title: "Map Internals"
description: "ค่าของ map คืออะไร การขยาย และทำไมลำดับสุ่ม"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

ค่าของ map คือ pointer ไปยัง hash table ใน runtime การ copy หรือส่ง map ไปจึงใช้ table เดียวกัน

## ส่ง map ไปให้ function

```go title="share.go"
func addItem(m map[string]int) { m["x"] = 1 } // ไม่ต้องใช้ pointer

m := make(map[string]int, 100) // บอกขนาดไว้ก่อน จะได้ไม่ต้องขยายบ่อย
addItem(m)
fmt.Println(m) // map[x:1]
```

## ข้างในเกิดอะไรขึ้น

- key ถูก hash ลง bucket ค้นหาเฉลี่ย O(1)
- พอ table แน่นเกินไปจะขยายและย้าย entry นี่คือเหตุผลที่ห้าม `&m[k]` เพราะ element ย้ายที่ได้
- การวนลูปเริ่มที่ตำแหน่งสุ่ม โค้ดจะได้พึ่งลำดับไม่ได้
- ตั้งแต่ Go 1.24 runtime ใช้โครงสร้างแบบ Swiss table เร็วขึ้น แต่จากฝั่งโค้ดทำงานเหมือนเดิม

:::caution[ข้อควรระวัง]
- map ไม่หดหลัง `delete` map ที่เคยมี key เป็นล้านจะยังถือ memory นั้นอยู่ ถ้าจำเป็นให้สร้างใหม่
:::
