---
title: "Implicit Interface Satisfaction"
description: "ไม่มี keyword implements แค่มี method ครบก็พอ"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

Go ไม่มี keyword `implements` type ตรงกับ interface ทันทีที่มี method ครบ แม้ interface จะเขียนทีหลังและอยู่คนละ package

## ตัวอย่าง

```go title="logger.go"
type Logger interface{ Log(msg string) }

type ConsoleLogger struct{ prefix string }

func (c ConsoleLogger) Log(msg string) { fmt.Println(c.prefix + msg) }

var _ Logger = ConsoleLogger{} // เช็คตอน compile (ไม่บังคับ)

func start(l Logger) { l.Log("started") }
start(ConsoleLogger{"[app] "}) // [app] started

var w io.Writer = &bytes.Buffer{} // bytes.Buffer ไม่เคยพูดถึง io.Writer เลย
```

## ทำไมถึงมีประโยชน์

ประกาศ interface เล็กๆ ครอบโค้ดที่ไม่ใช่ของเรา (client ของ third-party) แล้วสลับเป็นตัวปลอมตอน test ได้โดยไม่ต้องแตะโค้ดนั้น

:::caution[ข้อควรระวัง]
- พิมพ์ชื่อ method ผิด type จะไม่ตรงกับ interface แบบเงียบๆ บรรทัด `var _ I = T{}` ช่วยเปลี่ยนให้เป็น compile error
:::
