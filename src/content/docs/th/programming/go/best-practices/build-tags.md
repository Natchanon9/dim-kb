---
title: "Build Tags"
description: "ใส่ไฟล์เฉพาะบาง OS เวอร์ชัน หรือ tag ที่กำหนดเอง"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

บรรทัด `//go:build` ที่หัวไฟล์ตัดสินว่าไฟล์นั้นจะถูกรวมใน build หรือไม่

## ตัวอย่าง: สอง edition

```go title="edition_free.go"
//go:build !premium

package app

func Edition() string { return "free" }
```

```go title="edition_premium.go"
//go:build premium

package app

func Edition() string { return "premium" }
```

```sh
go build                 # Edition() → free
go build -tags premium   # Edition() → premium
```

## Constraint ที่เจอบ่อย

| Constraint | ความหมาย |
| --- | --- |
| `//go:build linux` | เฉพาะ Linux (`windows`, `darwin` …) |
| `//go:build amd64 \|\| arm64` | architecture |
| `//go:build go1.23` | Go เวอร์ชัน ≥ 1.23 |
| `//go:build integration` | tag ที่ตั้งเอง เช่น test ที่ช้า |
| `//go:build cgo` | เฉพาะเมื่อเปิด cgo |

ชื่อไฟล์ก็ใช้ได้: `file_linux.go`, `file_test.go`

:::caution[ข้อควรระวัง]
- บรรทัดนี้ต้องอยู่ก่อน `package` และตามด้วยบรรทัดว่าง ไม่อย่างนั้นเป็นแค่ comment
:::
