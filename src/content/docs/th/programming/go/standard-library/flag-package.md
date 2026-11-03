---
title: "flag Package"
description: "parse flag จาก command line"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

`flag` แปลง argument แบบ `-name=value` เป็นตัวแปรที่มี type พร้อมค่าเริ่มต้นและสร้างหน้า help `-h` ให้

## ตัวอย่าง

```go title="main.go"
port := flag.Int("port", 8080, "listen port")
debug := flag.Bool("debug", false, "enable debug logs")
flag.Parse()
fmt.Println(*port, *debug, flag.Args())
// ./app -port=9000 -debug extra → 9000 true [extra]
```

รูปแบบอื่น: `flag.StringVar(&cfg.Addr, "addr", ":8080", "...")`, `flag.Duration` และ `flag.NewFlagSet` สำหรับคำสั่งย่อย

## Flag vs env var

service ใน container ส่วนใหญ่อ่าน config จาก environment variable ส่วน flag เหมาะกับ CLI และการรันในเครื่อง หลายแอปรับทั้งสองแบบ โดยใช้ env เป็นค่าเริ่มต้นของ flag

:::caution[ข้อควรระวัง]
- flag ต้องมาก่อน argument ตามตำแหน่ง: `app -v file` ไม่ใช่ `app file -v`
- CLI ที่ซับซ้อนมีคำสั่งย่อยเยอะ นิยมใช้ `spf13/cobra`
:::
