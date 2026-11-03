---
title: "gob Encoding"
description: "รูปแบบ binary ของ Go สำหรับส่งข้อมูลระหว่างโปรแกรม Go"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`encoding/gob` เป็นรูปแบบ binary กะทัดรัดที่อธิบายตัวเองได้ ใช้ส่งค่า Go ระหว่างโปรแกรม Go ด้วยกัน

## ตัวอย่าง

```go title="gob.go"
type Point struct{ X, Y int }

var buf bytes.Buffer
gob.NewEncoder(&buf).Encode(Point{1, 2})
var p Point
gob.NewDecoder(&buf).Decode(&p)
fmt.Println(p) // {1 2}
```

## ใช้เมื่อไหร่

cache หรือไฟล์ที่อ่านโดยโค้ด Go ของเราเองเท่านั้น และ `net/rpc` ถ้าภาษาอื่นต้องอ่านด้วย ใช้ JSON หรือ Protocol Buffers

:::caution[ข้อควรระวัง]
- ค่าที่เป็น interface ต้อง `gob.Register(ConcreteType{})` ก่อน encode
- field ที่ไม่ export จะถูกข้าม
:::
