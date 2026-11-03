---
title: "Glossary"
description: "ศัพท์ที่ใช้ในเว็บนี้ ไทย / อังกฤษ ให้ตรงกัน"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

ศัพท์เหล่านี้คงเป็นภาษาอังกฤษในหน้าภาษาไทย คอลัมน์ความหมายคือคำนิยามบรรทัดเดียวที่ใช้ทั้งเว็บ

| ศัพท์ | ความหมาย |
| --- | --- |
| package | โฟลเดอร์ของไฟล์ `.go` ที่ใช้ชื่อเดียวกัน เป็นหน่วยของการมองเห็น |
| module | กลุ่ม package ที่มี `go.mod` เป็นหน่วยของการทำเวอร์ชัน |
| exported | ชื่อที่ขึ้นต้นด้วยตัวพิมพ์ใหญ่ package อื่นมองเห็นได้ |
| zero value | ค่าเริ่มต้นของ type เมื่อไม่ได้กำหนด (`0`, `""`, `nil`) |
| slice | มุมมองที่ขยายได้บน array: pointer, length, capacity |
| backing array | array ที่ slice ชี้อยู่ อาจถูกแชร์กัน |
| map | hash table จาก key ที่เทียบได้ไปยังค่า |
| rune | Unicode code point หนึ่งตัว (`int32`) |
| struct | กลุ่ม field ที่มีชื่อ |
| embedding | ใส่ type ใน struct โดยไม่ตั้งชื่อ field แล้ว member ถูก promote ขึ้นมา |
| method | function ที่มี receiver |
| receiver | ค่าที่ method ถูกเรียกบนมัน (ค่าหรือ pointer) |
| interface | ชุดของ method signature ที่ตรงกันได้โดยไม่ต้องประกาศ |
| type assertion | `x.(T)`: ดึงค่า concrete ออกจาก interface |
| generics / type parameter | `[T any]`: เขียนครั้งเดียวใช้ได้กับหลาย type |
| constraint | interface ที่จำกัดว่า type parameter รับ type ไหนได้ |
| pointer | address ของค่า `&x` เอา address ส่วน `*p` ตามไปที่ค่า |
| goroutine | function ที่ทำงานพร้อมกันแบบเบา จัดการโดย runtime ของ Go |
| channel | ท่อที่มี type สำหรับส่งค่าระหว่าง goroutine |
| data race | เข้าถึงพร้อมกันโดยไม่ synchronize และมีอย่างน้อยหนึ่งตัวเขียน |
| deadlock | ทุก goroutine ถูกบล็อกรอกันเอง |
| context | พกการยกเลิก deadline และค่าของ request |
| middleware | ตัวห่อ handler ที่ทำงานก่อน/หลังมัน |
| handler | function หรือ type ที่ตอบ HTTP request |
| sentinel error | ค่า error ที่ export ไว้ให้เทียบด้วย `errors.Is` |
| wrap | เติม context ให้ error โดยยังเก็บตัวเดิมไว้ (`%w`) |
| panic / recover | หยุด goroutine กะทันหัน / จับการหยุดนั้นใน deferred call |
| build tag | constraint `//go:build` ที่ตัดสินว่าไฟล์ไหนถูก compile |
| port / adapter | interface ที่ core ต้องการ / ตัว implement จริงของมัน |

:::caution[ข้อควรระวัง]
- ถ้าหน้าไทยต้องใช้ศัพท์ใหม่ ให้เพิ่มที่นี่ก่อน ทั้งสองภาษาจะได้ใช้คำเดียวกัน
:::
