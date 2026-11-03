---
title: "Internal Packages"
description: "ซ่อน package จากโค้ดนอก module ด้วย internal/"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

package ที่อยู่ใต้โฟลเดอร์ชื่อ `internal` ถูก import ได้เฉพาะโค้ดที่อยู่ใต้โฟลเดอร์แม่ของ `internal` นั้น compiler บังคับให้

## กติกา

```text
myapp/
├─ internal/format/    import ได้เฉพาะ myapp/...
├─ greet/              import internal/format ได้
└─ cmd/server/         import internal/format ได้
```

```go title="greet/greet.go"
import "github.com/me/myapp/internal/format"

func Hello(name string) string { return format.Title("hello, " + name) } // Hello, ...
```

module อื่นที่ import `github.com/me/myapp/internal/format` จะ compile ไม่ผ่าน: "use of internal package not allowed"

## ทำไมต้องใช้

เก็บรายละเอียดการ implement ไว้เป็นส่วนตัว จะได้แก้ได้โดยไม่ทำให้ใครพัง ใน service โค้ดส่วนใหญ่อยู่ใน `internal/`

:::caution[ข้อควรระวัง]
- `internal` จำกัดว่าใคร import ได้ ไม่ได้เปลี่ยนการมองเห็นชื่อ ชื่อที่ export ข้างในยังเป็น exported สำหรับผู้ที่ import ได้
:::
